package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// when a request comes in, the handler func is executed in a separate goroutine.
// Once the processing completes, the gooutine is terminated.
// This ensures that the server processes multiple requests concurrently.
// Runtime exceptions do not corrupt the server, but only on the goroutine.
func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8008"
	}

	mux := http.NewServeMux()
	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/v1/api", apiHandler)
	mux.HandleFunc("/v1/healthz", healthCheckHandler)
}

// w: object to write back the response; r: incoming request
// no return required
func apiHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "Hello World")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "Health shining like the moon!")
}

type requestLog struct {
	URL      string `json:"url"`
	Method   string `json:"method"`
	BodySize int64  `json:"content_length"`
	Protocol string `json:"protocol"`
}

func logRequest(r *http.Request) {
	l := requestLog{
		URL:      r.URL.String(),
		Method:   r.Method,
		BodySize: r.ContentLength,
		Protocol: r.Proto,
	}
	
	j, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}
	log.Println(string(j))
}
