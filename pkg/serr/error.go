package serr

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	ErrInternal       ErrorCode = "INTERNAL"
	ErrInvalidVoucher ErrorCode = "INVALID_VOUCHER"
	ErrReachLimit     ErrorCode = "REACH_LIMIT"
	ErrInvalidUser    ErrorCode = "INVALID_USER"
	ErrInvalidTime    ErrorCode = "INVALID_TIME"
	ErrInvalidInput   ErrorCode = "INVALID_INPUT"
)

type ServiceError struct {
	Method    string
	Cause     error
	Message   string
	ErrorCode ErrorCode
	Code      int
}

func (e ServiceError) Error() string {
	return fmt.Sprintf(
		"%s (%d) - %s: %s",
		e.Method, e.Code, e.Message, e.Cause,
	)
}

func ValidationErr(method, message string, code ErrorCode) error {
	return &ServiceError{
		Method:    method,
		Message:   message,
		Code:      http.StatusBadRequest,
		ErrorCode: code,
	}
}

func ServiceErr(method, message string, cause error, code int) error {
	return &ServiceError{
		Method:  method,
		Cause:   cause,
		Message: message,
		Code:    code,
	}
}

func DBError(method, repo string, cause error) error {
	err := &ServiceError{
		Method: fmt.Sprintf("%s.%s", repo, method),
		Cause:  cause,
	}
	switch {
	case errors.Is(cause, sql.ErrNoRows):
		err.Code = http.StatusNotFound
		err.Message = fmt.Sprintf("%s not found", repo)
	default:
		err.Code = http.StatusInternalServerError
		err.Message = fmt.Sprintf("could not perform action on %s", repo)
	}
	return err
}
