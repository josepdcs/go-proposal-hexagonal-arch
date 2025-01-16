package main

import (
	"log"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/di"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal("cannot load c: ", err)
	}

	server, err := di.InitializeAPI(c)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	} else {
		server.Start()
	}
}
