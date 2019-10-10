package main

import (
	"fmt"
	"net/http"
	"realtime-chat-go-react/pkg/websocket"
)

// "/ws" end point handler function
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}

func setupRoutes() {

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Simple server")
		serveWs(pool, w, r)
	})

}

func main() {
	setupRoutes()
	fmt.Println("Distributed app. Listning at localhost::8000")
	http.ListenAndServe(":8000", nil)

}
