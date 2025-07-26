package main

import (
	"clichat/client/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

func main() {
	var p *tea.Program
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:42069/ws", nil)

	if err != nil {
		fmt.Println("Websocket connection error :", err)
		os.Exit(1)
	}

	defer conn.Close()

	send := func(msg string) {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))

		if err != nil {
			p.Send(tui.IncomingMessage{Content: fmt.Sprintf("Error sending message :", err)})
		}
	}

	prog := tea.NewProgram(tui.InitialModel(send))
	p = prog

	go func() {
		for {
			_, message, err := conn.ReadMessage()

			if err != nil {
				p.Send(tui.IncomingMessage{Content: fmt.Sprintf("Disconnected")})
				return
			}
			p.Send(tui.IncomingMessage{Content: string(message)})
		}
	}()

	if _, err := prog.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
