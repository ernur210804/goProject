package app

import (
	"messengerApp/cmd/config"

	"database/sql"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/app/service"
)

type App struct {
	db             *sql.DB
	userRepo       repository.UserRepository
	friendRepo     repository.FriendRepository
	messageRepo    repository.MessageRepository
	profileRepo    repository.ProfileRepository
	authService    service.AuthService
	friendService  service.FriendService
	messageService service.MessageService
	profileService service.ProfileService
}

// Login implements service.AuthService.
func (a *App) Login(username string, password string) (string, error) {
	panic("unimplemented")
}

// Register implements service.AuthService.
func (a *App) Register(username string, password string) error {
	panic("unimplemented")
}

// RegisterUser implements service.AuthService.
func (a *App) RegisterUser(username string, password string) (*models.User, error) {
	panic("unimplemented")
}

func NewApp(db *sql.DB, cfg *config.Config) *App {
	// Create user repository
	userRepo := repository.NewUserRepository(db)

	// Create other repositories
	friendRepo := repository.NewFriendRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// Create service instances
	authService := service.NewAuthService(userRepo)
	friendService := service.NewFriendService(userRepo, friendRepo)
	messageService := service.NewMessageService(userRepo, friendRepo, messageRepo)
	profileService := service.NewProfileService(userRepo, profileRepo)

	return &App{
		db:             db,
		userRepo:       userRepo,
		friendRepo:     friendRepo,
		messageRepo:    messageRepo,
		profileRepo:    profileRepo,
		authService:    authService,
		friendService:  friendService,
		messageService: messageService,
		profileService: profileService,
	}
}

func (a *App) AuthService() service.AuthService {
	return a.authService
}

func (a *App) FriendService() service.FriendService {
	return a.friendService
}

func (a *App) MessageService() service.MessageService {
	return a.messageService
}

func (a *App) ProfileService() service.ProfileService {
	return a.profileService
}
