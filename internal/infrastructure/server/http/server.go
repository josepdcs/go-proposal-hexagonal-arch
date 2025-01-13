package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/thnkrn/go-gin-clean-arch/internal/infrastructure/server/middleware"

	_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/handler"
)

const (
	usersPath   = "users"
	usersPathID = usersPath + "/:id"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(user *handler.UserAPI) *Server {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", handler.Login)

	// Auth middleware
	api := engine.Group("/api", middleware.Authorization)

	api.GET(usersPath, user.FindAll)
	api.GET(usersPathID, user.FindByID)
	api.POST(usersPath, user.Create)
	api.PUT(usersPathID, user.Modify)
	api.DELETE(usersPathID, user.Delete)

	return &Server{engine: engine}
}

func (sh *Server) Start() {
	sh.engine.Run(":3000")
}
