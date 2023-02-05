package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8008"
	}
	
	http.HandleFunc("/v1/api", apiHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

// w: object to write back the response; r: incoming request
// no return required
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
