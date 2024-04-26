package repository

import (
	"database/sql"
	"messengerApp/internal/app/models"
)

type UserRepository interface {
	FindByID(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) FindByID(id int) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = ?"
	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	var user models.User
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
