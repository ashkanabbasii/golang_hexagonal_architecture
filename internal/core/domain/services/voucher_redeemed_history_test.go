package services_test

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"voucher/internal/core/domain/entity"
	"voucher/internal/core/domain/services"
	"voucher/internal/core/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVoucherRedeemedHistoryService_RecordRedemption(t *testing.T) {
	mockRedemptionPort := new(mocks.MockVoucherRedemptionPersistencePort)
	svc := services.NewVoucherRedeemedHistoryService(mockRedemptionPort)

	ctx := context.Background()
	voucherID := 1
	userID := "user123"
	tx := &sql.Tx{} // Assume tx is properly initialized for the test

	t.Run("Success", func(t *testing.T) {
		mockRedemptionPort.On("CreateRedeemedHistory", ctx, mock.AnythingOfType("*entity.VoucherRedemptionHistory"), tx).Return(nil).Once()

		err := svc.RecordRedemption(ctx, voucherID, userID, tx)
		assert.NoError(t, err)

		mockRedemptionPort.AssertExpectations(t)
	})

	t.Run("ValidationError", func(t *testing.T) {
		// Example: Missing userID (assuming Validate would fail on this)
		mockRedemptionPort.On("CreateRedeemedHistory", ctx, mock.AnythingOfType("*entity.VoucherRedemptionHistory"), tx).Return(nil).Maybe()

		err := svc.RecordRedemption(ctx, voucherID, "", tx)
		assert.Error(t, err)

		mockRedemptionPort.AssertNotCalled(t, "CreateRedeemedHistory")
	})
}

func TestVoucherRedeemedHistoryService_ListRedeemedHistoriesByCode(t *testing.T) {
	mockRedemptionPort := new(mocks.MockVoucherRedemptionPersistencePort)
	svc := services.NewVoucherRedeemedHistoryService(mockRedemptionPort)

	ctx := context.Background()
	code := "VOUCHERCODE"

	t.Run("Success", func(t *testing.T) {
		histories := []*entity.VoucherRedemptionHistory{
			{
				VoucherID:  1,
				UserID:     "user123",
				RedeemedAt: time.Now(),
			},
		}

		mockRedemptionPort.On("ListRedeemedHistoriesByCode", ctx, code).Return(histories, nil).Once()

		result, err := svc.ListRedeemedHistoriesByCode(ctx, code)
		assert.NoError(t, err)
		assert.Equal(t, histories, result)

		mockRedemptionPort.AssertExpectations(t)
	})

	t.Run("NoHistoriesFound", func(t *testing.T) {
		mockRedemptionPort.On("ListRedeemedHistoriesByCode", ctx, code).Return(nil, nil).Once()

		result, err := svc.ListRedeemedHistoriesByCode(ctx, code)
		assert.NoError(t, err)
		assert.Nil(t, result)

		mockRedemptionPort.AssertExpectations(t)
	})
}

func TestVoucherRedeemedHistoryService_ListRedeemedHistoriesByUser(t *testing.T) {
	mockRedemptionPort := new(mocks.MockVoucherRedemptionPersistencePort)
	svc := services.NewVoucherRedeemedHistoryService(mockRedemptionPort)

	ctx := context.Background()
	userID := "user123"

	t.Run("Success", func(t *testing.T) {
		histories := []*entity.VoucherRedemptionHistory{
			{
				VoucherID:  1,
				UserID:     userID,
				RedeemedAt: time.Now(),
			},
		}

		mockRedemptionPort.On("ListRedeemedHistoriesByUser", ctx, userID).Return(histories, nil).Once()

		result, err := svc.ListRedeemedHistoriesByUser(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, histories, result)

		mockRedemptionPort.AssertExpectations(t)
	})

	t.Run("NoHistoriesFound", func(t *testing.T) {
		mockRedemptionPort.On("ListRedeemedHistoriesByUser", ctx, userID).Return(nil, nil).Once()

		result, err := svc.ListRedeemedHistoriesByUser(ctx, userID)
		assert.NoError(t, err)
		assert.Nil(t, result)

		mockRedemptionPort.AssertExpectations(t)
	})
}

func TestVoucherRedeemedHistoryService_ListRedeemedHistoryUsage(t *testing.T) {
	mockRedemptionPort := new(mocks.MockVoucherRedemptionPersistencePort)
	svc := services.NewVoucherRedeemedHistoryService(mockRedemptionPort)

	ctx := context.Background()
	code := "VOUCHERCODE"
	userID := "user123"

	t.Run("Success", func(t *testing.T) {
		histories := []*entity.VoucherRedemptionHistory{
			{
				VoucherID:  1,
				UserID:     userID,
				RedeemedAt: time.Now(),
			},
		}

		mockRedemptionPort.On("ListRedeemedHistoryUsage", ctx, code, userID).Return(histories, nil).Once()

		result, err := svc.ListRedeemedHistoryUsage(ctx, code, userID)
		assert.NoError(t, err)
		assert.Equal(t, histories, result)

		mockRedemptionPort.AssertExpectations(t)
	})

	t.Run("NoHistoriesFound", func(t *testing.T) {
		mockRedemptionPort.On("ListRedeemedHistoryUsage", ctx, code, userID).Return(nil, nil).Once()

		result, err := svc.ListRedeemedHistoryUsage(ctx, code, userID)
		assert.NoError(t, err)
		assert.Nil(t, result)

		mockRedemptionPort.AssertExpectations(t)
	})

}
