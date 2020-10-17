package app

import (
	"gorm.io/gorm"
	"versla/util/logger"
)

type App struct {
	Logger *logger.Logger
	DB     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{
		Logger: logger,
		DB:     db,
	}
}
