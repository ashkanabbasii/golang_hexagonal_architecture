package entity

import (
	"time"
	"voucher/pkg/serr"
)

type (
	VoucherCode struct {
		ID           int          // Unique identifier for the voucher code
		Code         string       // The voucher code itself
		Amount       int          // Amount of credit
		State        VoucherState // State of the voucher (e.g., 'Available', 'Redeemed', 'Expired')
		UsageLimit   int          // Maximum number of redemptions allowed
		UserLimit    int          // Maximum number of user allowed
		CurrentUsage int          // Number of times the code has been Redeemed
		CreatedAt    time.Time    // When the code was created
		UpdatedAt    time.Time    // Last updated time
	}

	VoucherState string
)

const (
	Available VoucherState = "Available"
	Redeemed  VoucherState = "Redeemed"
	Expired   VoucherState = "Expired"
)

// Validate checks if the VoucherCode entity is valid
func (vc *VoucherCode) Validate() error {
	if vc.Code == "" {
		return serr.ValidationErr("VoucherCode.Validate",
			"voucher code cannot be empty",
			serr.ErrInvalidVoucher)
	}
	if vc.UsageLimit <= 0 {
		return serr.ValidationErr("VoucherCode.Validate",
			"usage limit must be greater than zero",
			serr.ErrInvalidVoucher)
	}
	if vc.UserLimit <= 0 {
		return serr.ValidationErr("VoucherCode.Validate",
			"user limit must be greater than zero",
			serr.ErrInvalidVoucher)
	}
	if vc.CurrentUsage < 0 {
		return serr.ValidationErr("VoucherCode.Validate",
			"current usage cannot be negative",
			serr.ErrInvalidVoucher)
	}
	if vc.State != Available && vc.State != Redeemed && vc.State != Expired {
		return serr.ValidationErr("VoucherCode.Validate",
			"invalid state",
			serr.ErrInvalidVoucher)
	}

	if vc.CurrentUsage >= vc.UsageLimit {
		return serr.ValidationErr("VoucherCode.Validate",
			"voucher usage is limited",
			serr.ErrReachLimit)
	}
	if vc.Amount <= 0 {
		return serr.ValidationErr("VoucherCode.Validate",
			"invalid amount",
			serr.ErrInvalidVoucher)
	}
	return nil
}
