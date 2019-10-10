package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined..."})
			}
			break
		case client := <-pool.Unregister:
			//pool.Clients[client] = false // ??
			delete(pool.Clients, client)
			fmt.Println("Size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "user left..."})
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending mesage to all client in the pool")
			for client := range pool.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
