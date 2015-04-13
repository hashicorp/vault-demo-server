package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/mattn/go-shellwords"
	"github.com/mitchellh/mapstructure"
)

// MessageHandler is the function called for messages on the websocket conn
type MessageHandler func(map[string]interface{}) error

func messageCLI(ws *websocket.Conn, vault *client) MessageHandler {
	return func(data map[string]interface{}) error {
		var req messageCLIRequest
		if err := mapstructure.WeakDecode(data, &req); err != nil {
			return err
		}

		args, err := shellwords.Parse(req.Command)
		if err != nil {
			return err
		}

		if args[0] != "vault" {
			return fmt.Errorf("only `vault` can be executed")
		}

		code, stdout, stderr := vault.CLI(args)
		return ws.WriteJSON(&messageCLIResponse{
			ExitCode: code,
			Stdout:   stdout,
			Stderr:   stderr,
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
