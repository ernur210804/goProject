// server/server.go
package server

import (
	"messengerApp/cmd/config"
	"messengerApp/internal/app/handlers"
	"messengerApp/internal/app/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	config      *config.ServerConfig
	authService service.AuthService
}

func NewServer(config *config.ServerConfig, authService service.AuthService) *Server {
	router := gin.Default()

	server := &Server{
		router:      router,
		config:      config,
		authService: authService,
	}

	server.setupRoutes()

	return server
}

func (s *Server) Run() {
	if s.config == nil || s.config.Port == "" {
		panic("Server configuration is missing or invalid")
	}
	s.router.Run(":" + s.config.Port)
}

func (s *Server) setupRoutes() {
	authHandler := handlers.NewAuthHandler(s.authService)

	s.router.POST("/login", authHandler.Login)
	s.router.POST("/register", authHandler.Register)

	// Other route setup for friend, message, profile handlers
}
