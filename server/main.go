package main

import (
	"bufio"
	"clichat/server/ws"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed Upgrading Protocol :", err)
		return
	}

	defer conn.Close()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()

			if err != nil {
				fmt.Println("Client Disconnected :", err)
				return
			}
			fmt.Println("Client:", string(msg))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("You:")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()

		if err := conn.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			fmt.Println("Write Error:", err)
			break
		}

	}
}
func main() {

	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Websocket server started on :42069")
	http.ListenAndServe(":42069", nil)
}
