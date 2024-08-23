package ports

import (
	"context"
	"database/sql"
	"voucher/internal/core/domain/entity"
)

type (
	// VoucherRedemptionPersistencePort defines the methods for interacting with voucher redeemed history data
	VoucherRedemptionPersistencePort interface {
		// CreateRedeemedHistory saves a new redemption record to the database
		CreateRedeemedHistory(ctx context.Context, history *entity.VoucherRedemptionHistory, tx *sql.Tx) error

		// ListRedeemedHistoriesByUser retrieves redemption records based on user id
		ListRedeemedHistoriesByUser(ctx context.Context, userID string) ([]*entity.VoucherRedemptionHistory, error)

		// ListRedeemedHistoriesByCode retrieves redemption records based on code
		ListRedeemedHistoriesByCode(ctx context.Context, code string) ([]*entity.VoucherRedemptionHistory, error)

		// ListRedeemedHistoryUsage retrieves redemption history by code and userID
		ListRedeemedHistoryUsage(ctx context.Context, code, userID string) ([]*entity.VoucherRedemptionHistory, error)
	}

	// VoucherRedeemedHistoryServicePort defines the methods for interacting with voucher redeemed history services
	VoucherRedeemedHistoryServicePort interface {
		// RecordRedemption create new voucher redeemed history entity
		RecordRedemption(ctx context.Context, voucherID int, userID string, tx *sql.Tx) error

		// ListRedeemedHistoriesByCode retrieves redemption history by code
		ListRedeemedHistoriesByCode(ctx context.Context, code string) ([]*entity.VoucherRedemptionHistory, error)

		// ListRedeemedHistoriesByUser retrieves redemption history by userID
		ListRedeemedHistoriesByUser(ctx context.Context, userID string) ([]*entity.VoucherRedemptionHistory, error)

		// ListRedeemedHistoryUsage retrieves redemption history by userID and code
		ListRedeemedHistoryUsage(ctx context.Context, code, userID string) ([]*entity.VoucherRedemptionHistory, error)
	}
)
