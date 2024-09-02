package dto

import (
	"time"
	"voucher/internal/core/domain/entity"
)

type (
	CreateVoucherRequest struct {
		Code        string    `json:"code" validate:"required"`
		Amount      int       `json:"amount" validate:"required,gt=0"`
		Description string    `json:"description" validate:"required"`
		UsageLimit  int       `json:"usage_limit" validate:"required,gt=0"`
		ExpiryDate  time.Time `json:"expiry_date" validate:"required"`
	}

	RedeemVoucherRequest struct {
		Code   string `json:"code" validate:"required"`
		UserID string `json:"user_id" validate:"required"`
	}

	ListRedeemVoucherByCodeRequest struct {
		Code string `json:"code" validate:"required"`
	}

	ListRedeemVoucherByUserIDRequest struct {
		UserID string `json:"user_id" validate:"required"`
	}

	VoucherRedemptionHistoryResponse struct {
		ID         int       `json:"id"`
		VoucherID  int       `json:"voucher_id"`
		Amount     int       `json:"amount"`
		RedeemedAt time.Time `json:"redeemed_at"`
		UserID     string    `json:"user_id"`
	}
)

func ToVoucherRedemptionHistoryEntity(entity *entity.VoucherRedemptionHistory) *VoucherRedemptionHistoryResponse {
	return &VoucherRedemptionHistoryResponse{
		ID:         entity.ID,
		VoucherID:  entity.VoucherID,
		Amount:     entity.Amount,
		RedeemedAt: entity.RedeemedAt,
		UserID:     entity.UserID,
	}
}
