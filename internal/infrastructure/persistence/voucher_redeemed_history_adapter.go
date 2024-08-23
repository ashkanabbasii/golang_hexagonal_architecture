package persistence

import (
	"context"
	"database/sql"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
	"voucher/internal/infrastructure/persistence/models"
	"voucher/pkg/serr"
)

type VoucherRedemptionPersistenceAdapter struct {
	db *sql.DB
}

// NewVoucherRedemptionPersistenceAdapter creates a new instance of VoucherRedemptionPersistenceAdapter.
func NewVoucherRedemptionPersistenceAdapter(db *sql.DB) ports.VoucherRedemptionPersistencePort {
	return &VoucherRedemptionPersistenceAdapter{db: db}
}

// CreateRedeemedHistory saves a new redemption record to the database.
func (r *VoucherRedemptionPersistenceAdapter) CreateRedeemedHistory(ctx context.Context, history *entity.VoucherRedemptionHistory, tx *sql.Tx) error {
	historyDB := models.ToVoucherRedemptionHistoryDB(history)

	query := `INSERT INTO voucher_redemption_history (voucher_id, amount, redeemed_at, user_id) VALUES ($1, $2, $3, $4)`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, historyDB.VoucherID, &history.Amount, historyDB.RedeemedAt, historyDB.UserID)
	} else {
		_, err = r.db.ExecContext(ctx, query, historyDB.VoucherID, &history.Amount, historyDB.RedeemedAt, historyDB.UserID)
	}
	if err != nil {
		return serr.DBError("CreateRedeemedHistory", "voucher_redeemed_history", err)
	}

	return nil
}

// ListRedeemedHistoriesByUser retrieves redemption records based on user id.
func (r *VoucherRedemptionPersistenceAdapter) ListRedeemedHistoriesByUser(ctx context.Context, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	query := `SELECT ` + models.VoucherRedemptionHistoryColumns + ` FROM voucher_redemption_history WHERE user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	historiesDB, err := models.ScanVoucherRedemptionHistories(rows)
	if err != nil {
		return nil, serr.DBError("ListRedeemedHistoriesByUser", "voucher_redeemed_history", err)
	}

	var histories []*entity.VoucherRedemptionHistory
	for _, historyDB := range historiesDB {
		histories = append(histories, historyDB.ToVoucherRedemptionHistoryEntity())
	}
	return histories, nil
}

// ListRedeemedHistoriesByCode retrieves redemption records based on voucher code.
func (r *VoucherRedemptionPersistenceAdapter) ListRedeemedHistoriesByCode(ctx context.Context, code string) ([]*entity.VoucherRedemptionHistory, error) {
	query := `SELECT ` + models.VoucherRedemptionHistoryColumns + ` FROM voucher_redemption_history WHERE code = $1`
	rows, err := r.db.QueryContext(ctx, query, code)
	if err != nil {
		return nil, serr.DBError("ListRedeemedHistoriesByCode", "voucher_redeemed_history", err)
	}
	defer rows.Close()

	historiesDB, err := models.ScanVoucherRedemptionHistories(rows)
	if err != nil {
		return nil, serr.DBError("ListRedeemedHistoriesByCode", "voucher_redeemed_history", err)
	}

	var histories []*entity.VoucherRedemptionHistory
	for _, historyDB := range historiesDB {
		histories = append(histories, historyDB.ToVoucherRedemptionHistoryEntity())
	}
	return histories, nil
}

// ListRedeemedHistoryUsage retrieves redemption records based on voucher code and userID.
func (r *VoucherRedemptionPersistenceAdapter) ListRedeemedHistoryUsage(ctx context.Context, code, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	query := `SELECT ` + models.VoucherRedemptionHistoryColumns + ` FROM voucher_redemption_history WHERE code = $1 AND user_id = $2`
	rows, err := r.db.QueryContext(ctx, query, code, userID)
	if err != nil {
		return nil, serr.DBError("ListRedeemedHistoryUsage", "voucher_redeemed_history", err)
	}
	defer rows.Close()

	historiesDB, err := models.ScanVoucherRedemptionHistories(rows)
	if err != nil {
		return nil, serr.DBError("ListRedeemedHistoryUsage", "voucher_redeemed_history", err)
	}

	var histories []*entity.VoucherRedemptionHistory
	for _, historyDB := range historiesDB {
		histories = append(histories, historyDB.ToVoucherRedemptionHistoryEntity())
	}
	return histories, nil
}
