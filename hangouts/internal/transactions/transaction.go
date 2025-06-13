package transactions

import (
	"log/slog"

	"gorm.io/gorm"
)

type Transactions struct {
	UserTransaction UserTransaction
}

func CreateTransactions(logger *slog.Logger, db *gorm.DB) *Transactions {
	userTransactions := CreateUserTransaction(db, logger)
	return &Transactions{UserTransaction: userTransactions}
}
