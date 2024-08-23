package services_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"
	"voucher/internal/core/domain/entity"
	"voucher/internal/core/domain/services"
	"voucher/internal/core/mocks"
	"voucher/pkg/serr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVoucherCodeService_CreateVoucher(t *testing.T) {
	mockPersistence := new(mocks.MockVoucherPersistencePort)
	svc := services.NewVoucherCodeService(mockPersistence)

	ctx := context.Background()
	code := "TESTVOUCHER"
	maxUsages := 10

	t.Run("Success", func(t *testing.T) {
		mockPersistence.On("GetVoucher", ctx, code).Return(nil, sql.ErrNoRows).Once()
		mockPersistence.On("CreateVoucher", ctx, mock.AnythingOfType("*entity.VoucherCode")).Return(nil).Once()

		err := svc.CreateVoucher(ctx, code, maxUsages)
		assert.NoError(t, err)

		mockPersistence.AssertExpectations(t)
	})

	t.Run("VoucherAlreadyExists", func(t *testing.T) {
		existingVoucher := &entity.VoucherCode{Code: code}
		expectedErr := serr.ValidationErr("VoucherCodeService.CreateVoucher", "voucher already exists", serr.ErrInvalidVoucher)

		mockPersistence.On("GetVoucher", ctx, code).Return(existingVoucher, expectedErr).Once()

		err := svc.CreateVoucher(ctx, code, maxUsages)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, expectedErr))

		mockPersistence.AssertExpectations(t)
	})
}

func TestVoucherCodeService_RedeemVoucher(t *testing.T) {
	mockPersistence := new(mocks.MockVoucherPersistencePort)
	svc := services.NewVoucherCodeService(mockPersistence)

	ctx := context.Background()
	code := "TESTVOUCHER"
	tx := &sql.Tx{} // Assume tx is properly initialized for the test

	t.Run("Success", func(t *testing.T) {
		voucher := &entity.VoucherCode{
			Code:         code,
			State:        entity.Available,
			Amount:       1000000,
			UsageLimit:   10,
			UserLimit:    1,
			CurrentUsage: 0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockPersistence.On("GetVoucherWithLock", ctx, code, tx).Return(voucher, nil).Once()
		mockPersistence.On("UpdateVoucher", ctx, voucher, tx).Return(nil).Once()

		redeemedVoucher, err := svc.RedeemVoucher(ctx, code, tx)
		assert.NoError(t, err)
		assert.Equal(t, voucher.Code, redeemedVoucher.Code)
		assert.Equal(t, 1, redeemedVoucher.CurrentUsage)

		mockPersistence.AssertExpectations(t)
	})

	t.Run("VoucherNotFound", func(t *testing.T) {
		expectedErr := serr.ValidationErr("VoucherCodeService.RedeemVoucher", "voucher not found", serr.ErrInvalidVoucher)
		mockPersistence.On("GetVoucherWithLock", ctx, code, tx).Return(nil, expectedErr).Once()
		voucher, err := svc.RedeemVoucher(ctx, code, tx)
		assert.Error(t, err)
		assert.Nil(t, voucher)
		assert.True(t, errors.Is(err, expectedErr))

		mockPersistence.AssertExpectations(t)
	})

	t.Run("UpdateVoucherFails", func(t *testing.T) {
		voucher := &entity.VoucherCode{
			Code:         code,
			State:        entity.Available,
			UsageLimit:   10,
			UserLimit:    1,
			Amount:       1000000,
			CurrentUsage: 0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mockPersistence.On("GetVoucherWithLock", ctx, code, tx).Return(voucher, nil).Once()
		mockPersistence.On("UpdateVoucher", ctx, voucher, tx).Return(errors.New("update failed")).Once()

		redeemedVoucher, err := svc.RedeemVoucher(ctx, code, tx)
		assert.Error(t, err)
		assert.Nil(t, redeemedVoucher)

		mockPersistence.AssertExpectations(t)
	})
}
