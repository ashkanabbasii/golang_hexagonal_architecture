package mocks

import (
	"context"
	"database/sql"
	"voucher/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// MockVoucherPersistencePort is a mock implementation of the VoucherPersistencePort interface
type MockVoucherPersistencePort struct {
	mock.Mock
}

func (m *MockVoucherPersistencePort) CreateVoucher(ctx context.Context, voucher *entity.VoucherCode) error {
	args := m.Called(ctx, voucher)
	return args.Error(0)
}

func (m *MockVoucherPersistencePort) GetVoucher(ctx context.Context, code string) (*entity.VoucherCode, error) {
	args := m.Called(ctx, code)
	if v := args.Get(0); v != nil {
		return v.(*entity.VoucherCode), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockVoucherPersistencePort) GetVoucherWithLock(ctx context.Context, code string, tx *sql.Tx) (*entity.VoucherCode, error) {
	args := m.Called(ctx, code, tx)
	if v := args.Get(0); v != nil {
		return v.(*entity.VoucherCode), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockVoucherPersistencePort) UpdateVoucher(ctx context.Context, voucher *entity.VoucherCode, tx *sql.Tx) error {
	args := m.Called(ctx, voucher, tx)
	return args.Error(0)
}
