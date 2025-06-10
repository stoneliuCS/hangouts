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
	srv, err := api.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8081", srv); err != nil {
		log.Fatal(err)
	}
}
