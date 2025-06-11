package controller

import (
	"hangouts/internal/database"
	"log/slog"
)

type Controller struct {
	db     database.Database
	logger *slog.Logger
}

func CreateController(db database.Database, logger *slog.Logger) Controller {
	return Controller{db: db, logger: logger}
}
