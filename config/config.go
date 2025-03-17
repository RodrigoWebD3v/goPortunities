package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("FAKE ERROR")
	//return nill
}

func GetLogger(p string) *Logger{
	logger = NewLogger(p)
	return logger
}