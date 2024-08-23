package dto

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate = validator.New()

// Validate method for CreateVoucherRequest
func (v *CreateVoucherRequest) Validate() error {
	return validate.Struct(v)
}

// Validate method for RedeemVoucherRequest
func (v *RedeemVoucherRequest) Validate() error {
	return validate.Struct(v)
}

// Validate method for ListRedeemVoucherByCodeRequest
func (v *ListRedeemVoucherByCodeRequest) Validate() error {
	return validate.Struct(v)
}

// Validate method for ListRedeemVoucherByUserIDRequest
func (v *ListRedeemVoucherByUserIDRequest) Validate() error {
	return validate.Struct(v)
}
