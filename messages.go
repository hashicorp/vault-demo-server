package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	shellwords "github.com/mattn/go-shellwords"
	"github.com/mitchellh/mapstructure"
)

// MessageHandler is the function called for messages on the websocket conn
type MessageHandler func(map[string]interface{}) error

func messageCLI(ws *websocket.Conn, vault *client) MessageHandler {
	return func(data map[string]interface{}) error {
		var req messageCLIRequest
		if err := mapstructure.WeakDecode(data, &req); err != nil {
			return ws.WriteJSON(&messageCLIResponse{
				ExitCode: 1,
				Stderr:   fmt.Sprintf("error decoding command: %s", err),
			})
		}

		args, err := shellwords.Parse(req.Command)
		if err != nil {
			return ws.WriteJSON(&messageCLIResponse{
				ExitCode: 1,
				Stderr:   fmt.Sprintf("error parsing command: %s", err),
			})
		}

		if len(args) == 0 || args[0] != "vault" {
			command := "<empty>"
			if len(args) > 0 {
				command = args[0]
			}

			return ws.WriteJSON(&messageCLIResponse{
				ExitCode: 127,
				Stderr:   fmt.Sprintf("invalid command: %s", command),
			})
		}

		log.Printf("[DEBUG] %s: executing: %v", ws.RemoteAddr(), args)
		code, stdout, stderr := vault.CLI(args[1:])
		return ws.WriteJSON(&messageCLIResponse{
			ExitCode: code,
			Stdout:   stdout,
			Stderr:   stderr,
		})
	}
}

func messagePing(ws *websocket.Conn, vault *client) MessageHandler {
	return func(data map[string]interface{}) error {
		return ws.WriteJSON(&messagePingResponse{
			Pong: true,
		})
	}
}

type messageCLIRequest struct {
	Command string `mapstructure:"command"`
}

type messageCLIResponse struct {
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
}

type messagePingResponse struct {
	Pong bool `json:"pong"`
}
