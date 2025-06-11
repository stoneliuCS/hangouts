package handler

import (
	"context"
	api "hangouts/gen"
	"hangouts/internal/controller"
	"hangouts/internal/database"
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
func NewHandler(db database.Database, logger *slog.Logger) api.Handler {
	return Handler{
		controller.CreateController(db, logger),
		logger,
	}
}

func (h Handler) NewError(ctx context.Context, err error) *api.ErrorSchemaStatusCode {
	return &api.ErrorSchemaStatusCode{StatusCode: 500, Response: api.ErrorSchema{Error: "Internal Server Error."}}
}

// Define a method on the Healthcheckservice method that pings the server.
func (h Handler) HealthcheckGet(ctx context.Context) (*api.HealthcheckGetOK, error) {
	return &api.HealthcheckGetOK{Message: api.OptHealthcheckGetOKMessage{Value: "OK", Set: true}}, nil
}

func (h Handler) Get(ctx context.Context) (api.GetOK, error) {
	html, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: openapiSpec,
	})
	if err != nil {
		return api.GetOK{}, err
	}
	return api.GetOK{Data: strings.NewReader(html)}, nil
}
