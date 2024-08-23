package main

import (
	"database/sql"
	"log"
	"voucher/cmd/clients"
	"voucher/internal/config"
	"voucher/internal/infrastructure/db"
	"voucher/internal/server"
)

func postgresDB() *sql.DB {
	psql, err := db.NewPostgres(
		config.DBName(), config.DBUser(), config.DBPassword(), config.DBHost(), config.DBPort(),
		config.DBMaxOpenConn(), config.DBMaxIdleConn(),
	)
	if err != nil {
		log.Fatalf("failed to initalize db: %v", err)
	}
	return psql
}

func externalClients() *clients.WalletClient {
	walletClient := clients.NewWallet(config.APIWalletInternal(), config.APIWalletExternal())

	return walletClient
}

func setupServer(s *server.Server, psql *sql.DB) {
	s.SetHealthFunc(healthFunc(psql)).SetupRoutes()
}
