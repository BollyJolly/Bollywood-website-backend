package websocket

type Room struct {
	Code string

	Clients map[*Client]bool

	Broadcast chan Message

	Register chan *Client

	Unregister chan *Client
}

func (r *Room) Run() {

	for {

		select {

		case client := <-r.Register:

			r.Clients[client] = true

		case client := <-r.Unregister:

			delete(r.Clients, client)

			close(client.Send)

		case message := <-r.Broadcast:

			for client := range r.Clients {

				client.Send <- message
			}
		}
	}
}