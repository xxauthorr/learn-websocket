package main

import (
	"gorilla-websocket/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("server starting at port 3000")

	_ = http.ListenAndServe(":3000", mux)
}
