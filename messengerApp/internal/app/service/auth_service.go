package service

import (
	"errors"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/utils"
)

// LoginRequest represents the JSON request payload for the login endpoint
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest represents the JSON request payload for the register endpoint
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthService represents the service layer interface for authentication-related operations
type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
	RegisterUser(username, password string) (*models.User, error) // Add RegisterUser method
}

// authService represents the concrete implementation of the AuthService interface
type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new instance of authService
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	// Authentication logic
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Register(username, password string) error {
	// Registration logic
	user := &models.User{Username: username, Password: password}
	err := s.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) RegisterUser(username, password string) (*models.User, error) {
	// Registration logic
	user := &models.User{Username: username, Password: password}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
