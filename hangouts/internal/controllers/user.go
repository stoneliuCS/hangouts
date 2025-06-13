package controllers

import (
	"context"
	api "hangouts/internal/api"
	"hangouts/internal/database/models"
	"hangouts/internal/services"
	"log/slog"
	"time"
)

type UserController interface {
	HandleCreateUser(ctx context.Context, req api.OptAPIV1UserPostReq) (api.APIV1UserPostRes, error)
}

type UserControllerImpl struct {
	logger      *slog.Logger
	userService services.UserService
}

func CreateUserController(logger *slog.Logger, userService services.UserService) UserController {
	return UserControllerImpl{logger: logger, userService: userService}
}

func (u UserControllerImpl) HandleCreateUser(ctx context.Context, req api.OptAPIV1UserPostReq) (api.APIV1UserPostRes, error) {
	// Deserialize
	user := &models.User{}
	user.Age = req.Value.GetAge()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Username = req.Value.GetUsername()
	user.FirstName = req.Value.GetFirstName()
	user.LastName = req.Value.GetLastName()
	res, err := u.userService.CreateUser(user)
	if err != nil {
		return &api.APIV1UserPostBadRequest{}, err
	}
	apiUser := api.APIV1UserPostCreated{FirstName: res.FirstName, LastName: res.LastName, ID: user.ID.String(), Age: user.Age}
	return &apiUser, nil
}
