package services

import (
	"context"
	"database/sql"
	"time"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
)

// VoucherRedeemedHistoryService provides domain logic related to voucher redemption history.
type VoucherRedeemedHistoryService struct {
	redemptionPort ports.VoucherRedemptionPersistencePort
}

// NewVoucherRedeemedHistoryService creates a new instance of VoucherRedeemedHistoryService.
func NewVoucherRedeemedHistoryService(redemptionPort ports.VoucherRedemptionPersistencePort) ports.VoucherRedeemedHistoryServicePort {
	return &VoucherRedeemedHistoryService{
		redemptionPort: redemptionPort,
	}
}

// RecordRedemption records a new voucher redemption in the history.
func (s *VoucherRedeemedHistoryService) RecordRedemption(ctx context.Context, voucherID int, userID string, tx *sql.Tx) error {
	history := &entity.VoucherRedemptionHistory{
		VoucherID:  voucherID,
		UserID:     userID,
		RedeemedAt: time.Now(),
	}
	err := history.Validate()
	if err != nil {
		return err
	}
	return s.redemptionPort.CreateRedeemedHistory(ctx, history, tx)
}

// ListRedeemedHistoriesByCode retrieves the redemption history for a specific voucher's code.
func (s *VoucherRedeemedHistoryService) ListRedeemedHistoriesByCode(ctx context.Context, code string) ([]*entity.VoucherRedemptionHistory, error) {
	return s.redemptionPort.ListRedeemedHistoriesByCode(ctx, code)
}

// ListRedeemedHistoriesByUser retrieves the redemption history for a specific user.
func (s *VoucherRedeemedHistoryService) ListRedeemedHistoriesByUser(ctx context.Context, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	return s.redemptionPort.ListRedeemedHistoriesByUser(ctx, userID)
}

// ListRedeemedHistoryUsage retrieves the redemption history usage by code and userID
func (s *VoucherRedeemedHistoryService) ListRedeemedHistoryUsage(ctx context.Context, code, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	return s.redemptionPort.ListRedeemedHistoryUsage(ctx, code, userID)
}
