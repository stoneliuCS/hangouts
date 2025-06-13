package database

import (
	"fmt"
	"hangouts/internal/database/models"
	"hangouts/internal/utils"
	"log/slog"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a database from the given environment config and available logger level.
func CreateDatabase(cfg utils.EnvConfig, logger *slog.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)
	db_creator := func() (*gorm.DB, error) {
		return gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: slogGorm.New(slogGorm.WithHandler(logger.Handler()),
				slogGorm.WithTraceAll(),
				slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug)),
			SkipDefaultTransaction: true,
		})
	}
	db := utils.SafeCall(db_creator)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
