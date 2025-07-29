package main

import (
	"bufio"
	"clichat/client/tui"
	"clichat/server/auth"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

var authManager = auth.NewAuthManager()

func StartChatClient(username string) {
	wsUrl := fmt.Sprintf("ws://localhost:42069/ws?user=%s", username)
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)

	if err != nil {
		log.Fatal("Websocket dial error :", err)
	}

	defer conn.Close()

	model := tui.InitialModel(func(msg string) {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Println("Error occured initiaising tui model:", err)
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
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to CliChat")
	fmt.Print("Do you want to (r)egister or (l)ogin :")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Print("Enter Username")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Enter Password:")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	switch choice {
	case "r":
		if err := authManager.Register(username, password); err != nil {
			fmt.Println("Error occured on Register", err)
			return
		}
	case "l":
		if err := authManager.Login(username, password); err != nil {
			fmt.Println("Error occured on Login", err)
			return
		}
	default:
		fmt.Println("Invalid choice")
		return
	}
	StartChatClient(username)
}
