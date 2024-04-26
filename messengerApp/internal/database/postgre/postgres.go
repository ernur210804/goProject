package database

import (
	"database/sql"
	"fmt"

	"messengerApp/cmd/config"

	_ "github.com/lib/pq"
)

func InitDatabase(cfg *config.DatabaseConfig) *sql.DB {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	return db
}
