# 🧋 CliChat

A lightweight CLI chat app built with [Go](https://golang.org), [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API), and the amazing [Bubbletea](https://github.com/charmbracelet/bubbletea) TUI framework.

> Minimal. Beautiful. Real-time.

---

## ✨ Features

- 📡 **Real-time messaging** using WebSockets
- 💬 **Interactive terminal UI** powered by Bubbletea & Lipgloss
- 🧠 **Simple architecture** with hub-based broadcast model
- 🔧 Built with idiomatic Go using `net/http` and `gorilla/websocket`

---

## 📁 Project Structure

```
chat-app/
│
├── server/ # WebSocket server (Go)
│   └── main.go
│   hub
│    └── hub.go
│   ws
│    └── handler.go
│
├── client/         # TUI chat client (Go)
│   ├── main.go
│   └── tui/
│       └── model.go
│
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### 1. Start the server

```bash
cd server
go run .
```

> Server runs on `localhost:42069/home/arya/Downloads/clichat.gif` by default

---

### 2. Run a client

Open a new terminal:

```bash
cd client
go run .
```

Open another terminal and run the client again to simulate another user.

---

## 🖼️ Demo

![hippo](https://drive.google.com/file/d/1ead7jw_xumWB9J7XOo0iSxme0nNygH47/view?usp=drive_link)
---

## 🧠 How It Works

- Each client connects to the server via WebSocket
- A central **Hub** manages all active connections
- Messages are sent to the hub and **broadcast to all clients**
- Each TUI client updates in real-time using Bubbletea's reactive model

---

## 📦 Dependencies

- [Go 1.20+](https://golang.org/)
- [`github.com/gorilla/websocket`](https://github.com/gorilla/websocket)
- [`github.com/charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea)
- [`github.com/charmbracelet/lipgloss`](https://github.com/charmbracelet/lipgloss)
- [`github.com/charmbracelet/bubbles`](https://github.com/charmbracelet/bubbles)

Install them using:

```bash
go mod tidy
```

---

## 💡 Ideas for Future

- 🔐 User authentication
- 🧑‍🤝‍🧑 Private messages / group chat
- 🧪 Add unit tests and e2e testing

---

## 🧊 Credits

- The awesome [Charm](https://github.com/charmbracelet) team for their terminal UI ecosystem
- Inspired by Go community WebSocket examples and TUI apps

---

## 📜 License

MIT — use it freely, improve it endlessly.
