package handlers

import (
	"messengerApp/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FriendRequest represents the JSON request payload for adding a friend
type FriendRequest struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
}

// MessageRequest represents the JSON request payload for sending a message
type MessageRequest struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
	Message  string `json:"message"`
}

type FriendHandler struct {
	friendService service.FriendService
}

func NewFriendHandler(friendService service.FriendService) *FriendHandler {
	return &FriendHandler{friendService: friendService}
}

func (h *FriendHandler) AddFriend(c *gin.Context) {
	var friendRequest FriendRequest
	if err := c.BindJSON(&friendRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Convert UserID from string to int
	userID, err := strconv.Atoi(friendRequest.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert FriendID from string to int
	friendID, err := strconv.Atoi(friendRequest.FriendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
		return
	}

	// Call the friend service to add the friend
	if err := h.friendService.AddFriend(userID, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend added successfully"})
}
func (h *FriendHandler) SendMessage(c *gin.Context) {
	var messageRequest MessageRequest
	if err := c.BindJSON(&messageRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Convert UserID from string to int
	senderID, err := strconv.Atoi(messageRequest.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Convert FriendID from string to int
	receiverID, err := strconv.Atoi(messageRequest.FriendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
		return
	}

	// Call the message service to send the message to friend
	if err := h.friendService.SendMessage(senderID, receiverID, messageRequest.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}
