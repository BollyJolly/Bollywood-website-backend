package websocket

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) GetRoom(code string) *Room {

	room, exists := h.Rooms[code]

	if !exists {

		room = &Room{
			Code: code,

			Clients: make(map[*Client]bool),

			Broadcast: make(chan Message),

			Register: make(chan *Client),

			Unregister: make(chan *Client),
		}

		h.Rooms[code] = room

		go room.Run()
	}

	return room
}