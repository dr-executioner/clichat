package main

import (
	"clichat/client/tui"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:42069/ws", nil)

	if err != nil {
		fmt.Println("Websocket connection error :", err)
		os.Exit(1)
	}

	defer conn.Close()

	model := tui.InitialModel(func(msg string) {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Println("Error occured:", err)
		}
	})

	p := tea.NewProgram(model, tea.WithAltScreen())

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read Error:", err)
				p.Quit()
				return
			}
			p.Send(tui.IncomingMessage{Content: string(msg)})
		}
	}()
	if _, err := p.Run(); err != nil {
		log.Println("Error Running TUI:", err)
		os.Exit(1)
	}

}
