package ws

import (
	"clichat/server/hub"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	Hub *hub.Hub
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Hub: hub.NewHub(),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (wss *WebSocketServer) WsHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade Error", err)
		return
	}
	go wss.Hub.HandleConnection(conn)
}
