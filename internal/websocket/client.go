package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn

	Room *Room

	UserID string

	Send chan Message
}

func (c *Client) ReadMessages() {

	defer func() {

		c.Room.Unregister <- c

		c.Conn.Close()

	}()

	for {

		var message Message

		err := c.Conn.ReadJSON(&message)

		if err != nil {
			break
		}

		message.UserID = c.UserID

		c.Room.Broadcast <- message
	}
}

func (c *Client) WriteMessages() {

	for message := range c.Send {

		err := c.Conn.WriteJSON(message)

		if err != nil {
			break
		}
	}
}