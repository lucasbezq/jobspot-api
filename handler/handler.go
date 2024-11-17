package handler

import (
	"github.com/lucasbezq/jobspot-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func IniializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
