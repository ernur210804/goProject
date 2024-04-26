// message_service.go

package service

import (
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type SendMessageRequest struct {
	UserID   int    `json:"user_id"`
	FriendID int    `json:"friend_id"`
	Message  string `json:"message"`
}
type MessageService interface {
	SendMessage(senderID, receiverID int, message string) error
	GetMessages(userID int) ([]*models.Message, error)
}

type messageService struct {
	userRepo    repository.UserRepository
	friendRepo  repository.FriendRepository
	messageRepo repository.MessageRepository
}

func NewMessageService(userRepo repository.UserRepository, friendRepo repository.FriendRepository, messageRepo repository.MessageRepository) MessageService {
	return &messageService{
		userRepo:    userRepo,
		friendRepo:  friendRepo,
		messageRepo: messageRepo,
	}
}

func (s *messageService) SendMessage(senderID, receiverID int, message string) error {
	// Simulate sending message by returning nil
	return nil
}

func (s *messageService) GetMessages(userID int) ([]*models.Message, error) {
	// Simulate getting messages by returning empty slice
	return []*models.Message{}, nil
}
