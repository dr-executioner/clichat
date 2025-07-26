package ws

import (
	"clichat/client/tui"
	"clichat/utils"
	"fmt"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

func StartWebSocketListener(p *tea.Program) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:42069/ws", nil)
	if err != nil {
		log.Printf("Websocket connection error: %v\n", err)
		p.Send(tui.IncomingMessage{Content: "Failed to connect to server"})
		return
	}

	go func() {
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()

			if err != nil {
				p.Send(tui.IncomingMessage{Content: "Disconnected"})
				return
			}
			p.Send(tui.IncomingMessage{Content: string(msg)})
		}

	}()
}

func WsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := utils.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error Upgradding to WebSocket Protocol", err)
		return
	}

	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("Error Occured:", err)
			break
		}
		fmt.Printf("Received: %s\n", msg)

		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Write Error:", err)
			return
		}
	}

}
