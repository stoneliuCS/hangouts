package services

import (
	"hangouts/internal/transactions"
	"log/slog"
)

// Services available in Hangouts
type Services struct {
	UserService UserService
}

// Creates all the services available.
func CreateServices(logger *slog.Logger, transactions *transactions.Transactions) *Services {
	userService := CreateUserService(logger, transactions.UserTransaction)
	return &Services{UserService: userService}
}
