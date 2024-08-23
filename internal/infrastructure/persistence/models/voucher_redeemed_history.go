package models

import (
	"database/sql"
	"time"
	"voucher/internal/core/domain/entity"
)

// VoucherRedemptionHistoryDB represents the database model for the voucher_redemption_history table
type VoucherRedemptionHistoryDB struct {
	ID         int
	VoucherID  int
	Amount     int
	RedeemedAt time.Time
	UserID     string
}

// VoucherRedemptionHistoryColumns Constants representing the column names for easier usage in queries.
const (
	VoucherRedemptionHistoryColumns = "id, amount, voucher_id, redeemed_at, user_id"
)

// ScanVoucherRedemptionHistory scans a database row into a VoucherRedemptionHistoryDB model.
func ScanVoucherRedemptionHistory(row *sql.Row) (*VoucherRedemptionHistoryDB, error) {
	var history VoucherRedemptionHistoryDB
	err := row.Scan(&history.ID, &history.Amount, &history.VoucherID, &history.RedeemedAt, &history.UserID)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// ScanVoucherRedemptionHistories scans multiple database rows into a slice of VoucherRedemptionHistoryDB models.
func ScanVoucherRedemptionHistories(rows *sql.Rows) ([]*VoucherRedemptionHistoryDB, error) {
	var histories []*VoucherRedemptionHistoryDB
	for rows.Next() {
		var history VoucherRedemptionHistoryDB
		err := rows.Scan(&history.ID, &history.Amount, &history.VoucherID, &history.RedeemedAt, &history.UserID)
		if err != nil {
			return nil, err
		}
		histories = append(histories, &history)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return histories, nil
}

// ToVoucherRedemptionHistoryEntity converts the database model to the core entity
func (v *VoucherRedemptionHistoryDB) ToVoucherRedemptionHistoryEntity() *entity.VoucherRedemptionHistory {
	return &entity.VoucherRedemptionHistory{
		ID:         v.ID,
		VoucherID:  v.VoucherID,
		Amount:     v.Amount,
		RedeemedAt: v.RedeemedAt,
		UserID:     v.UserID,
	}
}

// ToVoucherRedemptionHistoryDB converts the core entity to the database model
func ToVoucherRedemptionHistoryDB(v *entity.VoucherRedemptionHistory) *VoucherRedemptionHistoryDB {
	return &VoucherRedemptionHistoryDB{
		ID:         v.ID,
		VoucherID:  v.VoucherID,
		Amount:     v.Amount,
		RedeemedAt: v.RedeemedAt,
		UserID:     v.UserID,
	}
}
