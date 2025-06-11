package main

import (
	api "hangouts/gen"
	"hangouts/internal/controller"
	"hangouts/internal/handler"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	logger := slog.New(slog.Default().Handler())
	// Declare controller
	var c controller.Controller
	h := handler.NewHandler(c, logger)
	srv_func := func() (*api.Server, error) { return api.NewServer(h) }
	srv := SafeCall(srv_func)
	serve_func := func() error {
		return http.ListenAndServe(":8081", srv)
	}
	SafeCallErrorSupplier(serve_func)
}

func SafeCall[T any](fn func() (T, error)) T {
	val, err := fn()
	if err != nil {
		log.Fatal(err)
		panic("Fatal error.")
	}
	return val
}

func SafeCallErrorSupplier(fn func() error) {
	err := fn()
	if err != nil {
		log.Fatal(err)
		panic("Fatal error.")
	}
}
