package websocket

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


type Handler struct {
	hub *Hub
}

func NewHandler(
	hub *Hub,
) *Handler {

	return &Handler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) Connect(
	c *gin.Context,
) {

	fmt.Println("WebSocket request received")

	roomCode := c.Param("code")

	userID := c.GetString("userId")

	conn, err := upgrader.Upgrade(
		c.Writer,
		c.Request,
		nil,
	)

	if err != nil {
		fmt.Println("Upgrade failed:", err)
		return
	}

	fmt.Println("WebSocket connected")

	room := h.hub.GetRoom(roomCode)

	client := &Client{
		Conn:   conn,
		Room:   room,
		UserID: userID,
		Send:   make(chan Message),
	}

	room.Register <- client

	fmt.Println("Client registered")

	go client.WriteMessages()
	go client.ReadMessages()
}