package admin

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupServer() *Server {
	timeout, _ := time.ParseDuration("45s")
	return &Server{
		svc: &http.Server{
			Addr:         ":9090",
			Handler:      handler(),
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			IdleTimeout:  timeout,
		},
	}
}

// Server represents a holder around a net/http Server which
// is used for admin endpoints. (i.e. metrics, healthcheck)
type Server struct {
	svc *http.Server
}

// Start brings up the admin HTTP service. This call blocks.
func (s *Server) Listen() error {
	return s.svc.ListenAndServe()
}

// Shutdown unbinds the HTTP server.
func (s *Server) Shutdown() {
	s.svc.Shutdown(context.TODO())
}

func handler() http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r
}
