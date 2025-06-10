package handler

import (
	"context"
	api "hangouts/gen"
	"hangouts/internal/controller"
	"log/slog"
)

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
func (h Handler) APIV1HealthcheckGet(ctx context.Context) (*api.APIV1HealthcheckGetOK, error) {
	return &api.APIV1HealthcheckGetOK{Message: api.OptAPIV1HealthcheckGetOKMessage{"OK", true}}, nil
}

func (h Handler) NewError(ctx context.Context, err error) *api.ErrRespStatusCode {
	return &api.ErrRespStatusCode{StatusCode: 500, Response: api.ErrResp{Message: api.OptString{}}}
}
