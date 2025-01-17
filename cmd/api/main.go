package main

import (
	"log"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/di"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("cannot load c: ", err)
	}

	server, err := di.InitializeAPI(cfg.DB)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	} else {
		server.Start()
	}
}
