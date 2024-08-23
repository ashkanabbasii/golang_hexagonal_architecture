package mocks

import (
	"context"
	"database/sql"
	"voucher/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// MockVoucherRedemptionPersistencePort is a mock implementation of the VoucherRedemptionPersistencePort interface
type MockVoucherRedemptionPersistencePort struct {
	mock.Mock
}

func (m *MockVoucherRedemptionPersistencePort) CreateRedeemedHistory(ctx context.Context, history *entity.VoucherRedemptionHistory, tx *sql.Tx) error {
	args := m.Called(ctx, history, tx)
	return args.Error(0)
}

func (m *MockVoucherRedemptionPersistencePort) ListRedeemedHistoriesByUser(ctx context.Context, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	args := m.Called(ctx, userID)
	if v := args.Get(0); v != nil {
		return v.([]*entity.VoucherRedemptionHistory), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockVoucherRedemptionPersistencePort) ListRedeemedHistoriesByCode(ctx context.Context, code string) ([]*entity.VoucherRedemptionHistory, error) {
	args := m.Called(ctx, code)
	if v := args.Get(0); v != nil {
		return v.([]*entity.VoucherRedemptionHistory), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockVoucherRedemptionPersistencePort) ListRedeemedHistoryUsage(ctx context.Context, code, userID string) ([]*entity.VoucherRedemptionHistory, error) {
	args := m.Called(ctx, code, userID)
	if v := args.Get(0); v != nil {
		return v.([]*entity.VoucherRedemptionHistory), args.Error(1)
	}
	return nil, args.Error(1)
}
