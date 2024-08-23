package dto

import "time"

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
)
