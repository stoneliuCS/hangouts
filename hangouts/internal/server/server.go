package server

import (
	"fmt"
	api "hangouts/internal/api"
	"hangouts/internal/utils"
	"log/slog"
	"net/http"
)

// Runs the server api with the given handler.
func RunServer(handler api.Handler, cfg utils.EnvConfig, logger *slog.Logger) {
	// Configure security handler for bearer authentication
	securityHandler := createSecurityHandler(cfg.JWT_SECRET_KEY)
	// Create middleware for logging.
	opts := api.WithMiddleware(logging(logger))

	// Create server
	srvFunc := func() (*api.Server, error) { return api.NewServer(handler, securityHandler, opts) }
	srv := utils.SafeCall(srvFunc)
	addr := fmt.Sprintf(":%d", cfg.PORT)
	servFunc := func() error {
		logger.Info("Started server on localhost" + addr)
		return http.ListenAndServe(addr, srv)
	}

	// Run server indefinitely
	utils.SafeCallErrorSupplier(servFunc)
}
