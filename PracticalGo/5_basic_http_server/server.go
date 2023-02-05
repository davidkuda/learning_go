package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = "127.0.0.1:8008"
		
		log.Fatal(http.ListenAndServe(listenAddr, nil))
	}
}
