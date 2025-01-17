package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
