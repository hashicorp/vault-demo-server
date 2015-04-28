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
	// MessageTypeCLI is the message type for executing CLI commands
	MessageTypeCLI = "cli"

	// MessageTypePing is a ping that just keeps the connection alive
	// since JavaScript can't send a WebSocket ping.
	MessageTypePing = "ping"
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

	// CORS
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Upgrade the connection to a websocket connection
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERR] %s", err)
		http.Error(w, "Internal error", 500)
		return
	}

	// Note connection closed
	log.Printf("[INFO] %s: connection open", ws.RemoteAddr())

	// Create a new core
	vault, err := NewClient()
	if err != nil {
		log.Printf("[ERR] %s", err)
		http.Error(w, "Internal error", 500)
		return
	}
	defer vault.Close()

	// Handle each message
	for {
		var message wsMessage
		if err := ws.ReadJSON(&message); err != nil {
			if err != io.EOF {
				log.Printf("[ERR] %s", err)
				http.Error(w, "Bad request", 400)
			}

			break
		}

		// Depending on the type of the message do something else
		var handler MessageHandler
		switch message.Type {
		case MessageTypeCLI:
			handler = messageCLI(ws, vault)
		case MessageTypePing:
			// Do nothing
		default:
			http.Error(
				w,
				fmt.Sprintf("unknown message type: %s", message.Type),
				400)
		}

		if handler != nil {
			if err := handler(message.Data); err != nil {
				http.Error(w, fmt.Sprintf("error: %s", err), 400)
				break
			}
		}
	}

	// Note connection closed
	log.Printf("[INFO] %s: connection closed", ws.RemoteAddr())
}

type wsMessage struct {
	Type string
	Data map[string]interface{}
}
