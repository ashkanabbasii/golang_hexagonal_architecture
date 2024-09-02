package services

import (
	"context"
	"database/sql"
	"net/http"
	"voucher/internal/core/application/ports"
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
func (s *VoucherRedemptionHistoryApplicationService) ListRedeemedHistoriesByCode(ctx context.Context, request *dto.ListRedeemVoucherByCodeRequest) ([]*dto.VoucherRedemptionHistoryResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, serr.ValidationErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByCode",
			err.Error(), serr.ErrInvalidInput)
	}

	result, err := s.domainService.ListRedeemedHistoriesByCode(ctx, request.Code)
	if err != nil {
		return nil, serr.ServiceErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByCode",
			err.Error(), err, http.StatusInternalServerError)
	}

	response := make([]*dto.VoucherRedemptionHistoryResponse, 0, len(result))
	for _, v := range result {
		response = append(response, dto.ToVoucherRedemptionHistoryEntity(v))
	}

	return response, nil
}

// ListRedeemedHistoriesByUser retrieves the redemption history for a specific user.
func (s *VoucherRedemptionHistoryApplicationService) ListRedeemedHistoriesByUser(ctx context.Context, request *dto.ListRedeemVoucherByUserIDRequest) ([]*dto.VoucherRedemptionHistoryResponse, error) {
	err := request.Validate()
	if err != nil {
		return nil, serr.ValidationErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByUser",
			err.Error(), serr.ErrInvalidInput)
	}

	result, err := s.domainService.ListRedeemedHistoriesByUser(ctx, request.UserID)
	if err != nil {
		return nil, serr.ServiceErr("VoucherRedemptionHistoryApplicationService.ListRedeemedHistoriesByUser",
			err.Error(), err, http.StatusInternalServerError)
	}

	response := make([]*dto.VoucherRedemptionHistoryResponse, 0, len(result))
	for _, v := range result {
		response = append(response, dto.ToVoucherRedemptionHistoryEntity(v))
	}

	return response, nil

}
