package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mitchellh/osext"
)

var addr = flag.String("addr", ":8080", "http service address")
var token = flag.String("token", "", "token id (internal)")

var selfPath string

func main() {
	os.Setenv("VAULT_INTERACTIVE_DEMO_SERVER", "true")
	os.Setenv("VAULT_FORMAT", "table")
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()

	// If we have a token set, then we're in token handler mode. Do it!
	if *token != "" {
		return mainToken(*token, flag.Args())
	}

	// Set our own path for later
	var err error
	selfPath, err = osext.Executable()
	if err != nil {
		log.Printf("[FATAL] Error getting executable path:%s", err)
		return 1
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/socket", handleWebSocket)

	log.Println("[INFO] Starting server...")
	log.Printf("[INFO] Executable (selfPath) = %s", selfPath)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Printf("[FATAL] error starting: %s", err)
		return 1
	}

	return 0
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://vaultproject.io", 301)
}
