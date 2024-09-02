package api

import "voucher/internal/server"

func SetupVoucherCodeRoutes(s *server.Server, v *VoucherCodeHandler) {
	s.External.POST("vouchers", v.CreateVoucher)
	s.External.PATCH("vouchers/redeem", v.RedeemVoucher)
}

func SetupVoucherRedeemedHistoryRoutes(s *server.Server, v *VoucherRedeemedHistoryHandler) {
	s.External.GET("vouchers/:code/history", v.ListRedeemedHistoriesByCode)
	s.External.GET("vouchers/users/:user_id/history", v.ListRedeemedHistoriesByUser)
}
