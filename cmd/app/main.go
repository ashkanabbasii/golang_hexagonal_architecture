package main

import (
	"go.uber.org/fx"
	"voucher/internal/config"
	service2 "voucher/internal/core/application/services"
	"voucher/internal/core/domain/services"
	"voucher/internal/infrastructure/db"
	"voucher/internal/infrastructure/persistence"
	"voucher/internal/interfaces/api"
	"voucher/internal/server"
	"voucher/pkg/logger"
)

func main() {
	fx.New(
		fx.Provide(
			// external clients
			externalClients,

			// postgres
			postgresDB,

			// persistence
			persistence.NewPostgresVoucherCodePersistence,
			persistence.NewVoucherRedemptionPersistenceAdapter,

			// domain services
			services.NewVoucherCodeService,
			services.NewVoucherRedeemedHistoryService,

			// application services
			service2.NewVoucherApplicationService,
			service2.NewVoucherRedemptionHistoryApplicationService,

			// handlers
			api.NewVoucherCodeHandler,
			api.NewVoucherRedeemedHistoryHandler,

			// server
			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.Init,
			logger.SetupLogger,
			setupServer,
			db.Migrate,
			api.SetupVoucherCodeRoutes,
			api.SetupVoucherRedeemedHistoryRoutes,
			server.Run,
		),
	).Run()

}
