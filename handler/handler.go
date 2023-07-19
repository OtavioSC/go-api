package handler

import (
	"github.com/OtavioSC/go-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitiliazeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
