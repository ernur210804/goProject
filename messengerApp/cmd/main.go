// main.go
package main

import (
	"fmt"
	"messengerApp/cmd/config"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/app/server"
	"messengerApp/internal/app/service"
	database "messengerApp/internal/database/postgre"

	_ "github.com/lib/pq"
)

func main() {
	// Load application configuration
	dbConfig, serverConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		return
	}

	// Initialize the database connection
	db, err := database.InitDatabase(dbConfig)
	if err != nil {
		fmt.Printf("error initializing database: %v\n", err)
		return
	}

	// Pass the db.DB() to NewUserRepository
	userRepo := repository.NewUserRepository(db.DB())

	// Initialize the authentication service
	authService := service.NewAuthService(userRepo)

	// Start the HTTP server
	server := server.NewServer(serverConfig, authService)
	server.Run()
}
