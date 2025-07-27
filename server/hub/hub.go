package hub

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Sender  *websocket.Conn
	Content []byte
}

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Message
	lock      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		msg := <-h.broadcast

		h.lock.Lock()

		for conn := range h.clients {
			if conn == msg.Sender {
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, msg.Content); err != nil {
				fmt.Println("Error writing to the client:", err)
				conn.Close()
				delete(h.clients, conn)
			}
		}
		h.lock.Unlock()
	}
}

func (h *Hub) HandleConnection(conn *websocket.Conn) {
	h.lock.Lock()
	h.clients[conn] = true
	h.lock.Unlock()

	defer func() {
		h.lock.Lock()
		delete(h.clients, conn)
		h.lock.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client Disconnected or Error:", err)
			break
		}
		h.broadcast <- Message{Sender: conn, Content: msg}
	}
}
