package ports

import (
	"context"
	"database/sql"
	"voucher/internal/core/domain/entity"
)

type (
	// VoucherPersistencePort defines the methods for interacting with voucher data
	VoucherPersistencePort interface {
		// CreateVoucher saves a new voucher to the database
		CreateVoucher(ctx context.Context, voucher *entity.VoucherCode) error

		// GetVoucher retrieves a voucher by its code
		GetVoucher(ctx context.Context, code string) (*entity.VoucherCode, error)

		// GetVoucherWithLock retrieves a voucher by its code and lock the row
		GetVoucherWithLock(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error)

		// UpdateVoucher updates an existing voucher in the database
		UpdateVoucher(ctx context.Context, voucher *entity.VoucherCode, tx *sql.Tx) error
	}

	// VoucherCodeServicePort defines the methods for interacting with voucher code services
	VoucherCodeServicePort interface {
		// CreateVoucher create new voucher entity
		CreateVoucher(ctx context.Context, code string, maxUsages int) error

		// RedeemVoucher redeem voucher by code
		RedeemVoucher(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error)
	}
)
