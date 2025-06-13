package main

import (
	"context"
	"hangouts/cmd/server"
	"hangouts/internal/database"
	"hangouts/internal/handler"
	"hangouts/internal/services"
	"hangouts/internal/transactions"
	"hangouts/internal/utils"
	"log/slog"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	// Setup logger and environment variables.
	logger := slog.New(slog.Default().Handler())

	logger.Info("Loading environment variables...")
	env := loadEnv()

	logger.Info("Creating database from environment variables...")
	db := database.CreateDatabase(env, logger)

	logger.Info("Auto migrating database schemas...")
	database.AutoMigrate(db)

	logger.Info("Initializing transaction layer...")
	transactions := transactions.CreateTransactions(logger, db)

	logger.Info("Intializing service layer...")
	services := services.CreateServices(logger, transactions)

	logger.Info("Intializing handler layer...")

	h := handler.NewHandler(logger, services)

	logger.Info("Attaching handler and running server...")
	server.RunServer(h)
}

// Loads the environment variables as an EnvConfig
func loadEnv() utils.EnvConfig {
	var config utils.EnvConfig
	envFun := func() error { return envconfig.Process(context.Background(), &config) }
	utils.SafeCallErrorSupplier(envFun)
	return config
}
