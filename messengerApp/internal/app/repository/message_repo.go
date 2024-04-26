// message_repo.go

package repository

import (
	"database/sql"
	"messengerApp/internal/app/models"
)

type MessageRepository interface {
	SendMessage(senderID, receiverID int, message string) error
	GetMessages(userID int) ([]*models.Message, error)
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SendMessage(senderID, receiverID int, message string) error {
	// Implementation for sending a message
	return nil
}

func (r *messageRepository) GetMessages(userID int) ([]*models.Message, error) {
	// Implementation for getting messages for a user
	return nil, nil
}
