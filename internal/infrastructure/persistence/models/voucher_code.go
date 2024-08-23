package models

import (
	"database/sql"
	"time"
	"voucher/internal/core/domain/entity"
	"voucher/pkg/serr"
)

// VoucherCodeDB represents the database model for the voucher_codes table
type VoucherCodeDB struct {
	ID           int
	Code         string
	Amount       int
	State        string
	UsageLimit   int
	UserLimit    int
	CurrentUsage int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Constants for voucher_codes table column names
const (
	VoucherColumns     = "id, code, amount, state, usage_limit, user_limit, current_usage, created_at, updated_at"
	VoucherColumnsNoID = "code, amount, state, usage_limit, user_limit, current_usage, created_at, updated_at"
)

// ScanVoucherCode scans a row into a VoucherCodeDB model
func ScanVoucherCode(row *sql.Row) (*VoucherCodeDB, error) {
	voucher := &VoucherCodeDB{}
	err := row.Scan(
		&voucher.ID,
		&voucher.Code,
		&voucher.Amount,
		&voucher.State,
		&voucher.UsageLimit,
		&voucher.UserLimit,
		&voucher.CurrentUsage,
		&voucher.CreatedAt,
		&voucher.UpdatedAt,
	)
	if err != nil {
		return nil, serr.DBError("ScanVoucherCode", "voucher_code", err)
	}
	return voucher, nil
}

// ToVoucherCodeEntity converts the database model to the core entity
func (v *VoucherCodeDB) ToVoucherCodeEntity() *entity.VoucherCode {
	return &entity.VoucherCode{
		ID:           v.ID,
		Code:         v.Code,
		Amount:       v.Amount,
		State:        entity.VoucherState(v.State),
		UsageLimit:   v.UsageLimit,
		UserLimit:    v.UserLimit,
		CurrentUsage: v.CurrentUsage,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
}

// ToVoucherCodeDB converts the core entity to the database model
func ToVoucherCodeDB(v *entity.VoucherCode) *VoucherCodeDB {
	return &VoucherCodeDB{
		ID:           v.ID,
		Code:         v.Code,
		Amount:       v.Amount,
		State:        string(v.State),
		UsageLimit:   v.UsageLimit,
		UserLimit:    v.UserLimit,
		CurrentUsage: v.CurrentUsage,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}
}
