package services

import (
	"context"
	"database/sql"
	"net/http"
	"voucher/cmd/clients"
	"voucher/internal/core/application/ports"
	"voucher/internal/core/domain/entity"
	"voucher/internal/infrastructure/db"
	api "voucher/internal/interfaces/api/dto"
	"voucher/pkg/serr"
)

// VoucherApplicationService provides application logic for vouchers.
type VoucherApplicationService struct {
	voucherCodeDomainService       ports.VoucherCodeServicePort
	redemptionHistoryDomainService ports.VoucherRedeemedHistoryServicePort
	walletClient                   ports.WalletPort
}

// NewVoucherApplicationService creates a new instance of VoucherApplicationService.
func NewVoucherApplicationService(
	voucherDomainService ports.VoucherCodeServicePort,
	redemptionHistoryDomainService ports.VoucherRedeemedHistoryServicePort,
	walletClient ports.WalletPort,

) *VoucherApplicationService {
	return &VoucherApplicationService{
		voucherCodeDomainService:       voucherDomainService,
		redemptionHistoryDomainService: redemptionHistoryDomainService,
		walletClient:                   walletClient,
	}
}

// RedeemVoucher handles the redemption process of a voucher and interacts with the domain services.
func (s *VoucherApplicationService) RedeemVoucher(ctx context.Context, request *api.RedeemVoucherRequest) error {
	// Perform basic validation, e.g., check if the code is empty or invalid
	err := request.Validate()
	if err != nil {
		return serr.ValidationErr("VoucherApplicationService.RedeemVoucher",
			err.Error(), serr.ErrInvalidInput)
	}
	// Get usage of voucher by user
	usage, err := s.redemptionHistoryDomainService.ListRedeemedHistoryUsage(ctx, request.Code, request.UserID)
	if err != nil {
		return err
	}
	// todo: if we need dynamic limitation , we can check the user limit with our persistence
	// todo: but right now we just check already usage of voucher by user
	if len(usage) > 0 {
		return serr.ValidationErr("VoucherApplicationService.RedeemVoucher",
			"you've been reach to the limit", serr.ErrReachLimit)
	}
	var voucher *entity.VoucherCode
	// Call the domain services method to redeem the voucher by transaction
	err = db.Transaction(ctx, sql.LevelReadCommitted, func(tx *sql.Tx) error {
		// redeem a voucher
		voucher, err = s.voucherCodeDomainService.RedeemVoucher(ctx, request.Code, tx)
		if err != nil {
			return err
		}

		// record a redemption
		err = s.redemptionHistoryDomainService.RecordRedemption(ctx, voucher.ID, request.UserID, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return serr.ServiceErr("VoucherApplicationService.RedeemVoucher",
			err.Error(), err, http.StatusInternalServerError)
	}

	//todo: implement API call failure here
	// increase user wallet
	err = s.walletClient.IncreaseWalletBalance(ctx, &clients.UpdateWalletBalanceRequest{
		UserID: request.UserID,
		Amount: float64(voucher.Amount),
	})
	if err != nil {
		return err
	}

	return nil
}

// CreateVoucher create a new voucher
func (s *VoucherApplicationService) CreateVoucher(ctx context.Context, request *api.CreateVoucherRequest) error {
	err := request.Validate()
	if err != nil {
		return serr.ValidationErr("VoucherApplicationService.CreateVoucher",
			err.Error(), serr.ErrInvalidInput)
	}
	return s.voucherCodeDomainService.CreateVoucher(ctx, request.Code, request.UsageLimit)
}
