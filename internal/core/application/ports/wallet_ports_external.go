package ports

import (
	"context"
	"voucher/cmd/clients"
)

type WalletPort interface {
	DecreaseWalletBalance(ctx context.Context, req *clients.UpdateWalletBalanceRequest) error
	IncreaseWalletBalance(ctx context.Context, req *clients.UpdateWalletBalanceRequest) error
}
