package main

import (
	"log"

	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/config"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
