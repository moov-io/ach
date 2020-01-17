// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package admin implements an http.Server which can be used for operations
// and monitoring tools. It's designed to be shipped (and ran) inside
// an existing Go service.
package admin

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewServer returns an admin Server instance that handles Prometheus metrics
// and pprof requests.
// Callers can use ':0' to bind onto a random port and call BindAddr() for the address.
func NewServer(addr string) *Server {
	timeout, _ := time.ParseDuration("45s")

	var listener net.Listener
	if addr == ":0" {
		listener, _ = net.Listen("tcp", "127.0.0.1:0")
	} else {
		listener, _ = net.Listen("tcp", addr)
	}

	router := handler()
	svc := &Server{
		router:   router,
		listener: listener,
		svc: &http.Server{
			Addr:         listener.Addr().String(),
			Handler:      router,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
			IdleTimeout:  timeout,
		},
	}

	svc.AddHandler("/live", svc.livenessHandler())
	svc.AddHandler("/ready", svc.readinessHandler())
	return svc
}

// Server represents a holder around a net/http Server which
// is used for admin endpoints. (i.e. metrics, healthcheck)
type Server struct {
	router   *mux.Router
	svc      *http.Server
	listener net.Listener

	liveChecks  []*healthCheck
	readyChecks []*healthCheck
}

// BindAddr returns the server's bind address. This is in Go's format so :8080 is valid.
func (s *Server) BindAddr() string {
	if s == nil || s.svc == nil {
		return ""
	}
	return s.listener.Addr().String()
}

// Listen brings up the admin HTTP server. This call blocks until the server is Shutdown or panics.
func (s *Server) Listen() error {
	if s == nil || s.svc == nil || s.listener == nil {
		return nil
	}
	return s.svc.Serve(s.listener)
}

// Shutdown unbinds the HTTP server.
func (s *Server) Shutdown() {
	if s == nil || s.svc == nil {
		return
	}
	s.svc.Shutdown(context.TODO())
}

// AddHandler will append an http.HandlerFunc to the admin Server
func (s *Server) AddHandler(path string, hf http.HandlerFunc) {
	s.router.HandleFunc(path, hf)
}

// AddVersionHandler will append 'GET /version' route returning the provided version
func (s *Server) AddVersionHandler(version string) {
	s.AddHandler("/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(version))
	})
}

// profileEnabled returns if a given pprof handler should be
// enabled according to pprofHandlers and the PPROF_* environment
// variables.
//
// These profiles can be disabled by setting the appropriate PPROF_*
// environment variable. (i.e. PPROF_ALLOCS=no)
//
// An empty string, "yes", or "true" enables the profile. Any other
// value disables the profile.
func profileEnabled(name string) bool {
	k := fmt.Sprintf("PPROF_%s", strings.ToUpper(name))
	v := strings.ToLower(os.Getenv(k))
	return v == "" || v == "yes" || v == "true"
}

// Handler returns an http.Handler for the admin http service.
// This contains metrics and pprof handlers.
//
// No metrics specific to the handler are recorded.
//
// We only want to expose on the admin servlet because these
// profiles/dumps can contain sensitive info (raw memory).
func Handler() http.Handler {
	return handler()
}

func handler() *mux.Router {
	r := mux.NewRouter()

	// prometheus metrics
	r.Path("/metrics").Handler(promhttp.Handler())

	// always register index and cmdline handlers
	r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))

	if profileEnabled("profile") {
		r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	}
	if profileEnabled("symbol") {
		r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	}
	if profileEnabled("trace") {
		r.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	}

	// Register runtime/pprof handlers
	if profileEnabled("allocs") {
		r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	}
	if profileEnabled("block") {
		runtime.SetBlockProfileRate(1)
		r.Handle("/debug/pprof/block", pprof.Handler("block"))
	}
	if profileEnabled("goroutine") {
		r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	}
	if profileEnabled("heap") {
		r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	}
	if profileEnabled("mutex") {
		runtime.SetMutexProfileFraction(1)
		r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	}
	if profileEnabled("threadcreate") {
		r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	}

	return r
}
