package handler

import (
	"context"
	api "hangouts/internal/api"
	"hangouts/internal/database/models"
	"time"
)

func (h Handler) APIV1UserPost(ctx context.Context, req api.OptAPIV1UserPostReq) (api.APIV1UserPostRes, error) {
	user := &models.User{}
	user.Age = req.Value.GetAge()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Username = req.Value.GetUsername()
	user.FirstName = req.Value.GetFirstName()
	user.LastName = req.Value.GetLastName()
	res, err := h.services.UserService.CreateUser(user)
	if err != nil {
		return &api.APIV1UserPostBadRequest{}, err
	}
	apiUser := api.APIV1UserPostCreated{FirstName: res.FirstName, LastName: res.LastName, ID: user.ID.String(), Age: user.Age}
	return &apiUser, nil
}
