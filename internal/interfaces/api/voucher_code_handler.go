package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voucher/internal/core/application/services"
	"voucher/internal/interfaces/api/dto"
	"voucher/pkg/serr"
)

type VoucherCodeHandler struct {
	voucherAppService *services.VoucherApplicationService
}

func NewVoucherCodeHandler(voucherAppService *services.VoucherApplicationService) *VoucherCodeHandler {
	return &VoucherCodeHandler{voucherAppService: voucherAppService}
}

// CreateVoucher godoc
// @Summary      Create a new voucher
// @Description  Create a new voucher with code, description, usage limit, and expiry date.
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateVoucherRequest true "Create Voucher Request"
// @Success      201  {object}  nil
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /vouchers [post]
func (h *VoucherCodeHandler) CreateVoucher(c *gin.Context) {
	var req dto.CreateVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, serr.ValidationErr("handler.CreateVoucher",
			"invalid input", serr.ErrInvalidInput))
		return
	}

	err := h.voucherAppService.CreateVoucher(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// RedeemVoucher godoc
// @Summary      Redeem a voucher
// @Description  Redeem a voucher code for a user.
// @Tags         Vouchers
// @Accept       json
// @Produce      json
// @Param        request body dto.RedeemVoucherRequest true "Redeem Voucher Request"
// @Success      200  {object}  nil
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /vouchers/redeem [patch]
func (h *VoucherCodeHandler) RedeemVoucher(c *gin.Context) {
	var req dto.RedeemVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.voucherAppService.RedeemVoucher(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
