package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"sync"

	kv "github.com/hashicorp/vault-plugin-secrets-kv"
	"github.com/hashicorp/vault/builtin/logical/pki"
	"github.com/hashicorp/vault/builtin/logical/totp"
	"github.com/hashicorp/vault/builtin/logical/transit"
	"github.com/hashicorp/vault/command"
	vaulttoken "github.com/hashicorp/vault/command/token"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/physical/inmem"
	"github.com/hashicorp/vault/vault"
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
	inm, err := inmem.NewInmem(nil, nil)
	if err != nil {
		return nil, err
	}

	// Create the core, sealed and in-memory
	core, err := vault.NewCore(&vault.CoreConfig{
		// Heroku doesn't support mlock syscall
		DisableMlock: true,

		Physical: &Physical{
			Backend: inm,
			Limit:   64000,
		},

		LogicalBackends: map[string]logical.Factory{
			"transit": transit.Factory,
			"pki":     pki.Factory,
			"totp":    totp.Factory,
			"kv":      kv.Factory,
		},
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
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	runOpts := &command.RunOptions{
		TokenHelper: &vaulttoken.ExternalTokenHelper{
			BinaryPath: fmt.Sprintf("%s -token=%s", selfPath, v.id),
		},
		Address: fmt.Sprintf("http://%s", v.listener.Addr()),
		Stdout:  stdout,
		Stderr:  stderr,
	}

	log.Printf("[DEBUG] %s: running command: %v", runOpts.Address, raw)

	exitCode := command.RunCustom(raw, runOpts)
	return exitCode, stdout.String(), stderr.String()
}

func (v *client) Close() error {
	// Stop listening
	v.listener.Close()

	// Nil things to try to expedite GC even if something is lingering on this
	v.core = nil

	return nil
}
