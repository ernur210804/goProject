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
	authService service.AuthService // Change to authService
}

func NewServer(config *config.ServerConfig, authService service.AuthService) *Server { // Change the parameter name to authService
	router := gin.Default()

	server := &Server{
		router:      router,
		config:      config,
		authService: authService, // Change to authService
	}

	server.setupRoutes()

	return server
}

func (s *Server) Run() {
	s.router.Run(":" + s.config.Port)
}

func (s *Server) setupRoutes() {
	authHandler := handlers.NewAuthHandler(s.authService) // Pass authService directly

	s.router.POST("/login", authHandler.Login)
	s.router.POST("/register", authHandler.Register)

	// Other route setup for friend, message, profile handlers
}
