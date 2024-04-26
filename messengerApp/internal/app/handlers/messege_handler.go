package handlers

import (
	"messengerApp/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	var sendMessageRequest service.SendMessageRequest
	if err := c.BindJSON(&sendMessageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call MessageService's SendMessage method
	if err := h.messageService.SendMessage(sendMessageRequest.UserID, sendMessageRequest.FriendID, sendMessageRequest.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	userID := c.Param("userID") // Assuming you're passing the userID in the URL path
	// Convert userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	// Call MessageService's GetMessages method
	messages, err := h.messageService.GetMessages(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// Implement other methods as needed
