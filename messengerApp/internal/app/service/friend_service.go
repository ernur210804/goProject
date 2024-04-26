// friend_service.go

package service

import (
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type FriendService interface {
	AddFriend(userID, friendID int) error
	GetFriends(userID int) ([]*models.User, error)
	SendMessage(senderID, receiverID int, message string) error
}

type friendService struct {
	userRepo   repository.UserRepository
	friendRepo repository.FriendRepository
}

func NewFriendService(userRepo repository.UserRepository, friendRepo repository.FriendRepository) FriendService {
	return &friendService{
		userRepo:   userRepo,
		friendRepo: friendRepo,
	}
}

func (s *friendService) AddFriend(userID, friendID int) error {
	// Simulate adding friend by returning nil
	return nil
}

func (s *friendService) GetFriends(userID int) ([]*models.User, error) {
	// Simulate getting friends by returning empty slice
	return []*models.User{}, nil
}

func (s *friendService) SendMessage(senderID, receiverID int, message string) error {
	// Simulate sending message by returning nil
	return nil
}
