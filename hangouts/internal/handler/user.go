package handler

import (
	"context"
	api "hangouts/internal/api"
)

func (h Handler) APIV1UserPost(ctx context.Context, req api.OptAPIV1UserPostReq) (api.APIV1UserPostRes, error) {
	return h.controllers.UserController.HandleCreateUser(ctx, req)
}
