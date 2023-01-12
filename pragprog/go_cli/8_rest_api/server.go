package main

import (
	"net/http"
)
// A multiplexer (short: Mux) function maps incoming requests to the proper handlers
// based on the URL of the request.
// from Gerardi21, page 277:
// net/http provides a default Mux (DefaultServeMux), but registers routes globally.
// Other packages may register routes without your awareness, so it's a good
// security practice to use your own Mux.
// With your own Mux, you can add dependencies to it (e.g. file names, db conns).
// Finally, a custom Mux allows integrated testing.
// http.Handler interface -> type that responds to an HTTP request
func newMux(todoFile string) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", rootHandler)
	return m
}

func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content))
}
