package handlers

import (
	"net/http"

	"github.com/davidkuda/chap6/complex-server/config"
)

func Register(mux *http.ServeMux, cfg config.AppConfig) {
	mux.Handle("/healthz", &app{cfg: cfg, handler: healthCheckHandler})
	mux.Handle("/api", &app{cfg: cfg, handler: apiHandler})
	mux.Handle("/panic", &app{cfg: cfg, handler: panicHandler})
}
