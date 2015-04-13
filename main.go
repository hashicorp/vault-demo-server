package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/socket", handleWebSocket)

	log.Println("[INFO] Starting server...")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Printf("[FATAL] error starting: %s", err)
		return 1
	}

	return 0
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://vaultproject.io", 301)
}
