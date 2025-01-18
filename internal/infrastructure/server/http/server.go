package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/api/handler"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/middleware"

	_ "github.com/josepdcs/go-proposal-hexagonal-arch/cmd/api/docs"
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
	app.Get("/swagger/*", swagger.HandlerDefault)

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

func (sh *Server) Start() error {
	return sh.app.Listen(":8080")
}

func (sh *Server) Shutdown() error {
	return sh.app.Shutdown()
}

func (sh *Server) ShutdownWithTimeout(timeout time.Duration) error {
	return sh.app.ShutdownWithTimeout(timeout)
}
