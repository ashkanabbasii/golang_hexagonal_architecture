package services

import (
	"context"
	"database/sql"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
	"voucher/internal/interfaces/api/dto"
	"voucher/pkg/serr"
)

type VoucherRedemptionHistoryApplicationService struct {
	domainService ports.VoucherRedeemedHistoryServicePort
}

// NewVoucherRedemptionHistoryApplicationService creates a new instance of VoucherRedemptionHistoryApplicationService.
func NewVoucherRedemptionHistoryApplicationService(domainService ports.VoucherRedeemedHistoryServicePort) *VoucherRedemptionHistoryApplicationService {
	return &VoucherRedemptionHistoryApplicationService{
		domainService: domainService,
	}
}

// RecordRedemption records a new voucher redemption in the history.
func (s *VoucherRedemptionHistoryApplicationService) RecordRedemption(ctx context.Context, voucherID int, userID string, tx *sql.Tx) error {
	return s.domainService.RecordRedemption(ctx, voucherID, userID, tx)
}

// ListRedeemedHistoriesByCode retrieves the redemption history for a specific voucher's code.
func (s *VoucherRedemptionHistoryApplicationService) ListRedeemedHistoriesByCode(ctx context.Context, request *dto.ListRedeemVoucherByCodeRequest) ([]*entity.VoucherRedemptionHistory, error) {
	err := request.Validate()
	if err != nil {
		return nil, serr.ValidationErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByCode",
			err.Error(), serr.ErrInvalidInput)
	}
	return s.domainService.ListRedeemedHistoriesByCode(ctx, request.Code)
}

// ListRedeemedHistoriesByUser retrieves the redemption history for a specific user.
func (s *VoucherRedemptionHistoryApplicationService) ListRedeemedHistoriesByUser(ctx context.Context, request *dto.ListRedeemVoucherByUserIDRequest) ([]*entity.VoucherRedemptionHistory, error) {
	err := request.Validate()
	if err != nil {
		return nil, serr.ValidationErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByUser",
			err.Error(), serr.ErrInvalidInput)
	}
	return s.domainService.ListRedeemedHistoriesByUser(ctx, request.UserID)
}
