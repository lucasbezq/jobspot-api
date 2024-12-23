package main

import (
	"github.com/lucasbezq/jobspot-api/config"
	"github.com/lucasbezq/jobspot-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
