package persistence

import (
	"context"
	"database/sql"
	"errors"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
	"voucher/internal/infrastructure/persistence/models"
	"voucher/pkg/serr"
)

type PostgresVoucherPersistenceAdapter struct {
	db *sql.DB
}

func NewPostgresVoucherCodePersistence(db *sql.DB) ports.VoucherPersistencePort {
	return &PostgresVoucherPersistenceAdapter{db: db}
}

// CreateVoucher saves a new voucher to the database
func (r *PostgresVoucherPersistenceAdapter) CreateVoucher(ctx context.Context, voucher *entity.VoucherCode) error {
	dbModel := models.ToVoucherCodeDB(voucher)
	query := `
		INSERT INTO voucher_codes (` + models.VoucherColumnsNoID + `)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	err := r.db.QueryRowContext(ctx, query, dbModel.Code, dbModel.Amount, dbModel.State, dbModel.UsageLimit, dbModel.UserLimit, dbModel.CurrentUsage, dbModel.CreatedAt, dbModel.UpdatedAt).Scan(&dbModel.ID)
	if err != nil {
		return serr.DBError("CreateVoucher", "voucher_code", err)
	}
	// Update the entity ID after successful creation
	voucher.ID = dbModel.ID
	return nil
}

// GetVoucher retrieves a voucher by its code
func (r *PostgresVoucherPersistenceAdapter) GetVoucher(ctx context.Context, code string) (*entity.VoucherCode, error) {
	query := `
		SELECT ` + models.VoucherColumns + `
		FROM voucher_codes 
		WHERE code = $1`
	row := r.db.QueryRowContext(ctx, query, code)
	voucher, err := models.ScanVoucherCode(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, serr.DBError("GetVoucher", "voucher_code", err)
	}
	return voucher.ToVoucherCodeEntity(), nil
}

// GetVoucherWithLock retrieves a voucher by its code and locks the row
func (r *PostgresVoucherPersistenceAdapter) GetVoucherWithLock(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error) {
	query := `
		SELECT ` + models.VoucherColumns + `
		FROM voucher_codes 
		WHERE code = $1 
		FOR UPDATE`
	var row *sql.Row
	row = tx.QueryRowContext(ctx, query, code)
	voucher, err := models.ScanVoucherCode(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, serr.DBError("GetVoucherWithLock", "voucher_code", err)
	}
	return voucher.ToVoucherCodeEntity(), nil
}

// UpdateVoucher updates an existing voucher in the database
func (r *PostgresVoucherPersistenceAdapter) UpdateVoucher(ctx context.Context, voucher *entity.VoucherCode, tx *sql.Tx) error {
	dbModel := models.ToVoucherCodeDB(voucher)
	query := `
		UPDATE voucher_codes 
		SET code = $1, amount = $2, state = $3, usage_limit = $4, user_limit = $5, current_usage = $6, updated_at = $7
		WHERE id = $8`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, dbModel.Code, dbModel.Amount, dbModel.State, dbModel.UsageLimit, dbModel.UserLimit, dbModel.CurrentUsage, dbModel.UpdatedAt, dbModel.ID)
	} else {
		_, err = r.db.ExecContext(ctx, query, dbModel.Code, dbModel.Amount, dbModel.State, dbModel.UsageLimit, dbModel.UserLimit, dbModel.CurrentUsage, dbModel.UpdatedAt, dbModel.ID)
	}
	if err != nil {
		return serr.DBError("UpdateVoucher", "voucher_code", err)
	}

	return nil
}
