package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voucher/internal/core/application/services"
	"voucher/internal/interfaces/api/dto"
)

type VoucherRedeemedHistoryHandler struct {
	voucherRedeemedHistoryAppService *services.VoucherRedemptionHistoryApplicationService
}

func NewVoucherRedeemedHistoryHandler(voucherAppService *services.VoucherRedemptionHistoryApplicationService) *VoucherRedeemedHistoryHandler {
	return &VoucherRedeemedHistoryHandler{voucherRedeemedHistoryAppService: voucherAppService}
}

// ListRedeemedHistoriesByCode godoc
// @Summary      List redeemed voucher histories by code
// @Description  Get a list of voucher redemption histories filtered by voucher code.
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        code path string true "Voucher Code"
// @Success      200  {array}   entity.VoucherRedemptionHistory
// @Failure      500  {object}  Error
// @Router       /vouchers/{code}/history [get]
func (h *VoucherRedeemedHistoryHandler) ListRedeemedHistoriesByCode(c *gin.Context) {
	var req dto.ListRedeemVoucherByCodeRequest
	req.Code = c.Param("code")
	result, err := h.voucherRedeemedHistoryAppService.ListRedeemedHistoriesByCode(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ListRedeemedHistoriesByUser godoc
// @Summary      List redeemed voucher histories by user ID
// @Description  Get a list of voucher redemption histories filtered by user ID.
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        user_id path string true "User ID"
// @Success      200  {array}   entity.VoucherRedemptionHistory
// @Failure      500  {object}  Error
// @Router       /vouchers/users/{user_id}/history [get]
func (h *VoucherRedeemedHistoryHandler) ListRedeemedHistoriesByUser(c *gin.Context) {
	var req dto.ListRedeemVoucherByUserIDRequest
	req.UserID = c.Param("user_id")
	result, err := h.voucherRedeemedHistoryAppService.ListRedeemedHistoriesByUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
