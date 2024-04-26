package main

import (
	"messengerApp/cmd/config"
	"messengerApp/internal/app"
	"messengerApp/internal/app/server"
	database "messengerApp/internal/database/postgre"
)

func main() {
	// Load application configuration
	cfg := config.LoadConfig()

	// Initialize the database connection
	db := database.InitDatabase(cfg.Database)

	// Initialize the application
	application := app.NewApp(db, cfg)

	// Initialize the authentication service
	// Start the HTTP server
	server := server.NewServer(cfg.Server, application)
	server.Run()
}
