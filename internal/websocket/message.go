package websocket

type Message struct {
	Type    string `json:"type"`
	UserID  string `json:"userId,omitempty"`
	User    string `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
	RoomCode string `json:"roomCode,omitempty"`

	Number int    `json:"number,omitempty"`
	URL    string `json:"url,omitempty"`
}