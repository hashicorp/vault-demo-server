package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// The message types that can be sent in the JSON
const (
	MessageTypeCLI = "cli"
)

// wsUpgrader is the upgrader we're using.
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(*http.Request) bool { return true },
}

// handleWebSocket handles websocket requests
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// Upgrade the connection to a websocket connection
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERR] %s", err)
		http.Error(w, "Internal error", 500)
		return
	}

	// Create a new core
	vault, err := NewClient()
	if err != nil {
		log.Printf("[ERR] %s", err)
		http.Error(w, "Internal error", 500)
	}
	defer vault.Close()

	// Handle each message
	for {
		var message wsMessage
		if err := ws.ReadJSON(&message); err != nil {
			if err == io.EOF {
				break
			}

			log.Printf("[ERR] %s", err)
			http.Error(w, "Bad request", 400)
			continue
		}

		// Depending on the type of the message do something else
		var handler MessageHandler
		switch message.Type {
		case MessageTypeCLI:
			handler = messageCLI(ws, vault)
		default:
			http.Error(
				w,
				fmt.Sprintf("unknown message type: %s", message.Type),
				400)
		}

		if handler != nil {
			if err := handler(message.Data); err != nil {
				http.Error(w, fmt.Sprintf("error: %s", err), 400)
			}
		}
	}
}

type wsMessage struct {
	Type string
	Data map[string]interface{}
}
