package main

import (
	"context"
	api "hangouts/gen"
	"hangouts/internal/database"
	"hangouts/internal/handler"
	"hangouts/internal/utils"
	"log/slog"
	"net/http"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	// Setup logger and environment variables.
	logger := slog.New(slog.Default().Handler())

	logger.Info("Loading environment variables...")
	env := loadEnv()

	// Create the database
	logger.Info("Creating database from environment variables.")
	db := database.CreateGormPostgresDatabase(env, logger)

	logger.Info("Creating handler from OpenAPI Specification...")
	h := createHandler(logger, db)

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

// Creates the handlers to be used for the server.
func createHandler(logger *slog.Logger, db database.Database) api.Handler {
	// Declare controller
	h := handler.NewHandler(db, logger)
	return h
}

// Loads the environment variables as an EnvConfig
func loadEnv() utils.EnvConfig {
	var config utils.EnvConfig
	envFun := func() error { return envconfig.Process(context.Background(), &config) }
	utils.SafeCallErrorSupplier(envFun)
	return config
}
