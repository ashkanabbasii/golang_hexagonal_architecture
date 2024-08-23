package services

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
	"voucher/pkg/serr"
)

// VoucherCodeService provides domain logic related to vouchers
type VoucherCodeService struct {
	persistencePort ports.VoucherPersistencePort
}

// NewVoucherCodeService creates a new instance of VoucherCodeService
func NewVoucherCodeService(persistencePort ports.VoucherPersistencePort) ports.VoucherCodeServicePort {
	return &VoucherCodeService{persistencePort: persistencePort}
}

// CreateVoucher creates a new voucher
func (s *VoucherCodeService) CreateVoucher(ctx context.Context, code string, maxUsages int) error {
	// Check if voucher already exists
	existingVoucher, err := s.persistencePort.GetVoucher(ctx, code)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if existingVoucher != nil {
		return serr.ValidationErr("VoucherCodeService.CreateVoucher",
			"voucher already exists", serr.ErrInvalidVoucher)
	}

	// Create and save new voucher
	voucher := &entity.VoucherCode{
		Code:         code,
		State:        entity.Available,
		UsageLimit:   maxUsages,
		CurrentUsage: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return s.persistencePort.CreateVoucher(ctx, voucher)
}

// RedeemVoucher handles the redemption process of a voucher
func (s *VoucherCodeService) RedeemVoucher(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error) {
	// todo: here add acquiring application level lock something like redis distribute lock
	// todo: for preventing parallel race condition
	// Retrieve the voucher
	voucher, err := s.persistencePort.GetVoucherWithLock(ctx, code, tx)
	if err != nil {
		return nil, err
	}
	// check voucher is not nil
	if voucher == nil {
		return nil, serr.ValidationErr("VoucherCodeService.RedeemVoucher",
			"voucher not found", serr.ErrInvalidVoucher)
	}

	// Update voucher usage count
	voucher.CurrentUsage++

	// validate voucher entity before update
	err = voucher.Validate()
	if err != nil {
		return nil, err
	}

	if err := s.persistencePort.UpdateVoucher(ctx, voucher, tx); err != nil {
		return nil, err
	}

	return voucher, nil
}
