package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(log logger.Interface) *gorm.DB {
	cfg := &gorm.Config{
		Logger: log,
	}
	db, err := gorm.Open(sqlite.Open("cocktails.db"), cfg)
	if err != nil {
		panic(err)
	}
	return db
}
