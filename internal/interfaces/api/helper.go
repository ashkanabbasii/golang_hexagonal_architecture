package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"voucher/pkg/serr"
)

type (
	Error struct {
		Message string         `json:"message"`
		Code    serr.ErrorCode `json:"code"`
		TraceID string         `json:"trace_id"`
	}

	Response[D any, S bool, M string] struct {
		Data    D `json:"data"`
		Status  S `json:"status"`
		Message M `json:"message"`
	}
)

func handleError(ctx *gin.Context, err error) {
	var serviceError *serr.ServiceError
	switch {
	case errors.As(err, &serviceError):
		var e *serr.ServiceError
		errors.As(err, &e)
		l := log.Error().Str("method", e.Method).Str("code", string(e.ErrorCode))
		if e.Cause != nil {
			l.Err(e.Cause)
		}
		l.Msg(e.Message)
		ctx.AbortWithStatusJSON(
			e.Code,
			Error{Code: e.ErrorCode, Message: e.Message},
		)
		return
	default:
		log.Error().Err(err).Msg("unknown error")
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			Error{Code: serr.ErrInternal, Message: "internal server error"},
		)
		return
	}
}
