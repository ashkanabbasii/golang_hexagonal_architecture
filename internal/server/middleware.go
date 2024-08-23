package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"voucher/internal/config"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     config.CORSAllowedOrigins(),
		AllowMethods:     config.CORSAllowedMethods(),
		AllowHeaders:     config.CORSAllowedHeaders(),
		AllowCredentials: config.CORSAllowCredentials(),
		MaxAge:           12 * time.Hour,
	})
}
