package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

type Environment string

const (
	LOCAL Environment = "dev"
	BETA  Environment = "beta"
	PROD  Environment = "prod"
)

func Env() Environment {
	return Environment(viper.GetString("env"))
}

func ServiceName() string {
	return viper.GetString("domain.name")
}

func ServerInternalPort() int {
	return viper.GetInt("server.ports.internal")
}

func ServerExternalPort() int {
	return viper.GetInt("server.ports.external")
}

func ServerDebug() bool {
	return viper.GetBool("server.debug")
}

func ServerAddress() string {
	return viper.GetString("server.address")
}

func DBName() string {
	return viper.GetString("db.postgres.name")
}

func DBPassword() string {
	return viper.GetString("db.postgres.password")
}

func DBUser() string {
	return viper.GetString("db.postgres.user")
}

func DBPort() string {
	return viper.GetString("db.postgres.port")
}

func DBHost() string {
	return viper.GetString("db.postgres.host")
}

func DBMaxIdleConn() int {
	return viper.GetInt("db.postgres.maxIdleConn")
}

func DBMaxOpenConn() int {
	return viper.GetInt("db.postgres.maxOpenConn")
}

func DBMigrationsPath() string {
	return viper.GetString("db.postgres.migrationsPath")
}

func APIWalletExternal() string {
	return viper.GetString("api.wallet.external")
}

func APIWalletInternal() string {
	return viper.GetString("api.wallet.internal")
}

func CORSAllowedOrigins() []string {
	return viper.GetStringSlice("app.cors.allow-origins")
}

func CORSAllowedMethods() []string {
	return viper.GetStringSlice("app.cors.allow-methods")
}

func CORSAllowedHeaders() []string {
	return viper.GetStringSlice("app.cors.allow-headers")
}

func CORSAllowCredentials() bool {
	return viper.GetBool("app.cors.allow-credentials")
}

func Init() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "dev"))
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./internal/config/env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func getEnv(key, fallback string) string {
	log.Info().Msg("getting environment")
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
