package test

import "gorm.io/gorm/logger"

func NewTestLogger() logger.Interface {
	return logger.Default.LogMode(logger.Info)
}
