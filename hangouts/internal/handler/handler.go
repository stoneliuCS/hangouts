package handler

import (
	"context"
	api "hangouts/internal/api"
	"hangouts/internal/services"
	"log/slog"
	"strings"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
)

var openapiSpec string = "../openapi.json"

// Handles incoming API requests
type Handler struct {
	services *services.Services
	logger      *slog.Logger // event logger
}

// Creates a new handler for all defined API endpoints
func NewHandler(logger *slog.Logger, services *services.Services) api.Handler {
	return Handler{
		services,
		logger,
	}
}

func (h Handler) NewError(ctx context.Context, err error) *api.ErrRespStatusCode {
	return &api.ErrRespStatusCode{StatusCode: 401, Response: api.ErrResp{Error: "Unauthorized."}}
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
