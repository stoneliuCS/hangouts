package services

import (
	"hangouts/internal/database/models"
	"hangouts/internal/transactions"
	"log/slog"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *models.User) (*models.User, error)
}

type UserServiceImpl struct {
	userTransaction transactions.UserTransaction
	logger          *slog.Logger
}

// CreateUser implements UserService.
func (u UserServiceImpl) CreateUser(req *models.User) (*models.User, error) {
	req.ID = uuid.New()
	return u.userTransaction.CreateUser(req)
}

func CreateUserService(logger *slog.Logger, userTransaction transactions.UserTransaction) UserService {
	return UserServiceImpl{userTransaction: userTransaction, logger: logger}
}
