package main

import (
	api "hangouts/gen"
	"hangouts/internal/controller"
	"hangouts/internal/handler"
	"hangouts/internal/utils"
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	// Declare controller
	var c controller.Controller
	h := handler.NewHandler(c, logger)
	srv_func := func() (*api.Server, error) { return api.NewServer(h) }
	srv := utils.SafeCall(srv_func)
	serve_func := func() error {
		return http.ListenAndServe(":8081", srv)
	}
	utils.SafeCallErrorSupplier(serve_func)
}
