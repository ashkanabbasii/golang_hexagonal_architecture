package entity

import (
	"time"
	"voucher/pkg/serr"
)

// VoucherRedemptionHistory represents the redemption record of a voucher code
type VoucherRedemptionHistory struct {
	ID         int       // Unique identifier for the redemption record
	VoucherID  int       // ID of the Redeemed voucher
	Amount     int       // Amount of voucher
	RedeemedAt time.Time // When the voucher was Redeemed
	UserID     string    // ID of the user who Redeemed the voucher
}

func (vrh *VoucherRedemptionHistory) Validate() error {
	if vrh.VoucherID <= 0 {
		return serr.ValidationErr("VoucherRedemptionHistory.Validate",
			"invalid voucher ID",
			serr.ErrInvalidVoucher)
	}
	if vrh.UserID == "" {
		return serr.ValidationErr("VoucherRedemptionHistory.Validate",
			"invalid user ID",
			serr.ErrInvalidUser)
	}
	if vrh.RedeemedAt.IsZero() {
		return serr.ValidationErr("VoucherRedemptionHistory.Validate",
			"Redeemed at timestamp cannot be zero",
			serr.ErrInvalidTime)
	}
	return nil
}
