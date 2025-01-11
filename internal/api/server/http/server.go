package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/server/middleware"

	_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
	"github.com/thnkrn/go-gin-clean-arch/internal/api/handler"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(user *handler.User) *Server {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// Request JWT
	engine.POST("/login", handler.Login)

	// Auth middleware
	api := engine.Group("/api", middleware.Authorization)

	api.GET("users", user.FindAll)
	api.GET("users/:id", user.FindByID)
	api.POST("users", user.Save)
	api.DELETE("users/:id", user.Delete)

	return &Server{engine: engine}
}

func (sh *Server) Start() {
	sh.engine.Run(":3000")
}
