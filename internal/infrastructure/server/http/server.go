package http

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/josepdcs/go-proposal-hexagonal-arch/cmd/api/docs"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/api/handler"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/middleware"
)

const (
	usersPath   = "users"
	usersPathID = usersPath + "/:id"
)

type Server struct {
	app *fiber.App
}

func NewServer(user *handler.UserAPI) *Server {
	app := fiber.New()

	// Swagger docs
	//app.Get("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	app.Post("/login", handler.Login)

	// Auth middleware
	api := app.Group("/api", middleware.Authorization)

	api.Get(usersPath, user.FindAll)
	api.Get(usersPathID, user.FindByID)
	api.Post(usersPath, user.Create)
	api.Put(usersPathID, user.Modify)
	api.Delete(usersPathID, user.Delete)

	return &Server{app: app}
}

func (sh *Server) Start() {
	sh.app.Listen(":8080")
}
