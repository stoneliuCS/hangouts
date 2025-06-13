package transactions

import (
	"hangouts/internal/database/models"
	"log/slog"

	"gorm.io/gorm"
)

type UserTransaction interface {
	CreateUser(req *models.User) (*models.User, error)
}

type UserTransactionImpl struct {
	db     *gorm.DB
	logger *slog.Logger
}

func (u UserTransactionImpl) CreateUser(req *models.User) (*models.User, error) {
	res := u.db.Create(req)
	if res.Error != nil {
		return nil, res.Error
	}
	return req, nil
}

func CreateUserTransaction(db *gorm.DB, logger *slog.Logger) UserTransaction {
	return UserTransactionImpl{db: db, logger: logger}
}
