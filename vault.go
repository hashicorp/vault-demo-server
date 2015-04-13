package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"

	vaultcli "github.com/hashicorp/vault/cli"
	vaultcommand "github.com/hashicorp/vault/command"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/physical"
	"github.com/hashicorp/vault/vault"
	"github.com/mitchellh/cli"
)

// client represents a single vault instance: a single websocket connection.
// When this is closed, the Vault on the other side will be cleared.
type client struct {
	id       string
	core     *vault.Core
	listener net.Listener

	// internal
	lock sync.Mutex
}

// NewClient creates a new in-memory vault.
func NewClient() (*client, error) {
	// Create the core, sealed and in-memory
	core, err := vault.NewCore(&vault.CoreConfig{
		Physical: physical.NewInmem(),
	})
	if err != nil {
		return nil, err
	}

	// Create the HTTP server on a random local port, and start listening
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	go http.Serve(ln, vaulthttp.Handler(core))

	return &client{
		id:       strconv.FormatInt(int64(rand.Int31n(math.MaxInt32)), 10),
		core:     core,
		listener: ln,
	}, nil
}

// CLI is used to execute a CLI command on the local Vault. The address
// to the in-memory Vault will be prepended to the args.
func (v *client) CLI(raw []string) (int, string, string) {
	var stdout, stderr bytes.Buffer

	// Build our CLI commands
	commands := vaultcli.Commands(&vaultcommand.Meta{
		Ui: &cli.BasicUi{
			Writer:      &stdout,
			ErrorWriter: &stderr,
		},

		ForceAddress: v.listener.Addr().String(),
		ForceConfig: &vaultcommand.Config{
			TokenHelper: fmt.Sprintf("%s -token=%s", selfPath, v.id),
		},
	})

	exitCode := vaultcli.RunCustom(raw, commands)
	return exitCode, stdout.String(), stderr.String()
}

func (v *client) Close() error {
	// Stop listening
	v.listener.Close()

	// Nil things to try to expedite GC even if something is lingering on this
	v.core = nil

	return nil
}
