package handlers

import (
	"fmt"
	"net/http"

	"github.com/davidkuda/chap6/complex-server/config"
)

type app struct {
	cfg     config.AppConfig
	handler func(w http.ResponseWriter, r *http.Request, cfg config.AppConfig)
}

func (a app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler(w, r, a.cfg)
}

func apiHandler(w http.ResponseWriter, r *http.Request, cfg config.AppConfig) {
	fmt.Fprintf(w, "HelloWorld")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, cfg config.AppConfig) {
	if r.Method != http.MethodGet {
		cfg.Logger.Printf("error: InvalidRequest; path: %s; method: %s", r.URL.Path, r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "ok")
}

func panicHandler(w http.ResponseWriter, r *http.Request, cfg config.AppConfig) {
	panic("I panicked")
}
