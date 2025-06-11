package handler

import (
	"context"
	api "hangouts/gen"
	"hangouts/internal/controller"
	"hangouts/internal/utils"
	"log/slog"
	"strings"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
)

var openapiSpec string = "../openapi.json"

// Handles incoming API requests
type Handler struct {
	controller controller.Controller // executes business logic
	logger     *slog.Logger          // event logger
}

// Creates a new handler for all defined API endpoints
func NewHandler(controller controller.Controller, logger *slog.Logger) api.Handler {
	return Handler{
		controller,
		logger,
	}
}

// Define a method on the Healthcheckservice method that pings the server.
func (h Handler) APIV1HealthcheckGet(ctx context.Context) (api.APIV1HealthcheckGetRes, error) {
	return &api.APIV1HealthcheckGetOK{Message: api.OptAPIV1HealthcheckGetOKMessage{Value: "OK", Set: true}}, nil
}

func (h Handler) Get(ctx context.Context) (api.GetOK, error) {
	scalar_func := func() (string, error) {
		return scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: openapiSpec,
		})
	}
	html := utils.SafeCall(scalar_func)
	return api.GetOK{Data: strings.NewReader(html)}, nil
}
