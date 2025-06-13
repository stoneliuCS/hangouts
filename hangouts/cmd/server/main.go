package main

import (
	"context"
	api "hangouts/internal/api"
	"hangouts/internal/database"
	"hangouts/internal/handler"
	"hangouts/internal/services"
	"hangouts/internal/transactions"
	"hangouts/internal/utils"
	"log/slog"
	"net/http"

	"github.com/sethvargo/go-envconfig"
	"gorm.io/gorm"
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
	transactions := createTransactions(logger, db)

	logger.Info("Intializing service layer...")
	services := createServices(logger, transactions)

	logger.Info("Intializing handler layer...")
	h := createHandler(logger, services)

	logger.Info("Attaching handler and running server...")
	runServer(h)
}

// Runs the server api with the given handler.
func runServer(handler api.Handler) {
	srv_func := func() (*api.Server, error) { return api.NewServer(handler) }
	srv := utils.SafeCall(srv_func)
	serve_func := func() error {
		return http.ListenAndServe(":8081", srv)
	}
	utils.SafeCallErrorSupplier(serve_func)
}

func createServices(logger *slog.Logger, transactions *transactions.Transactions) *services.Services {
	userService := services.CreateUserService(logger, transactions.UserTransaction)
	return &services.Services{UserService: userService}
}

func createTransactions(logger *slog.Logger, db *gorm.DB) *transactions.Transactions {
	userTransactions := transactions.CreateUserTransaction(db, logger)
	return &transactions.Transactions{UserTransaction: userTransactions}
}

// Creates the handlers to be used for the server.
func createHandler(logger *slog.Logger, services *services.Services) api.Handler {
	// Declare controller
	h := handler.NewHandler(logger, services)
	return h
}

// Loads the environment variables as an EnvConfig
func loadEnv() utils.EnvConfig {
	var config utils.EnvConfig
	envFun := func() error { return envconfig.Process(context.Background(), &config) }
	utils.SafeCallErrorSupplier(envFun)
	return config
}
