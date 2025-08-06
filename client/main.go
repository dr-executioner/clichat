package main

import (
	"bufio"
	"clichat/client/tui"
	"clichat/server/auth"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

type OutGoingMessage struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

var authManager, err = auth.NewAuthManager("root:qwertyroot@tcp(127.0.0.1:3306)/clichatdb")

func StartChatClient(username string) {
	wsUrl := fmt.Sprintf("ws://localhost:42069/ws?user=%s", username)
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)

	if err != nil {
		log.Fatal("Websocket dial error :", err)
	}

	defer conn.Close()

	model := tui.InitialModel(func(msg string) {
		payload :=
			OutGoingMessage{
				From:    username,
				Content: msg,
			}
		jsonBytes, _ := json.Marshal(payload)
		if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
			log.Println("Error occured initiaising tui model:", err)
		}
	})

	p := tea.NewProgram(model, tea.WithAltScreen())

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				p.Send(tui.IncomingMessage{From: "System", Content: "Disconnected"})
				return
			}

			var outMsg OutGoingMessage
			if err := json.Unmarshal(msg, &outMsg); err != nil {
				p.Send(tui.IncomingMessage{From: "Unknown", Content: string(msg)})
				continue
			}
			p.Send(tui.IncomingMessage{From: outMsg.From, Content: outMsg.Content})
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
