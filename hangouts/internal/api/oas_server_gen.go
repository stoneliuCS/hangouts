// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIV1UserPost implements POST /api/v1/user operation.
	//
	// POST /api/v1/user
	APIV1UserPost(ctx context.Context, req OptAPIV1UserPostReq) (APIV1UserPostRes, error)
	// Get implements GET / operation.
	//
	// API documentation.
	//
	// GET /
	Get(ctx context.Context) (GetOK, error)
	// HealthcheckGet implements GET /healthcheck operation.
	//
	// GET /healthcheck
	HealthcheckGet(ctx context.Context) (*HealthcheckGetOK, error)
	// NewError creates *ErrRespStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrRespStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
