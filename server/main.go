package main

import (
	"clichat/server/hub"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {

	hub := hub.NewHub()

	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println("Server side upgrade websocket error", err)
			return
		}

		hub.HandleConnection(conn)
	})

	fmt.Println("Server Started on :42069")
	http.ListenAndServe(":42069", nil)
}
