package main

import (
	"github.com/OtavioSC/go-api/config"
	"github.com/OtavioSC/go-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}