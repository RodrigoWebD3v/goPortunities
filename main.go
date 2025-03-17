package main

import (
	"github.com/RodrigoWebD3v/goPortunities/config"
	"github.com/RodrigoWebD3v/goPortunities/router"
)

//Initialize configs

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
