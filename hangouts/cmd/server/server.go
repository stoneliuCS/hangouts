package server

import (
	api "hangouts/internal/api"
	"hangouts/internal/utils"
	"net/http"
)

// Runs the server api with the given handler.
func RunServer(handler api.Handler) {
	srv_func := func() (*api.Server, error) { return api.NewServer(handler) }
	srv := utils.SafeCall(srv_func)
	serve_func := func() error {
		return http.ListenAndServe(":8081", srv)
	}
	utils.SafeCallErrorSupplier(serve_func)
}
