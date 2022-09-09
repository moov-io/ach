package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/moov-io/base/admin"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/http/bind"
	"github.com/moov-io/base/strx"

	"github.com/moov-io/ach"

	"github.com/gorilla/mux"
)

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ACH"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ACH"), "Admin HTTP listen address")

	flagBasePath = flag.String("base-path", "/", "Base path to serve HTTP routes and webui from")
)

func main() {
	flag.Parse()

	log.Printf("Starting moov-io/ach webui version %s", ach.Version)

	// Channel for errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Start Admin server (with Prometheus metrics)
	adminServer := admin.NewServer(*adminAddr)
	adminServer.AddVersionHandler(ach.Version) // Setup 'GET /version'
	go func() {
		log.Printf("listening on %s", adminServer.BindAddr())
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			log.Print(err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	// Setup business HTTP routes
	router := mux.NewRouter().PathPrefix(*flagBasePath).Subrouter()
	addPingRoute(router)

	// Register our assets route
	assetsPath := strx.Or(os.Getenv("ASSETS_PATH"), filepath.Join("cmd", "webui", "assets"))
	log.Printf("serving assets from %s", assetsPath)
	addAssetsPath(router, assetsPath)

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			log.Print(err)
		}
	}

	// Start business logic HTTP server
	go func() {
		if certFile, keyFile := os.Getenv("HTTPS_CERT_FILE"), os.Getenv("HTTPS_KEY_FILE"); certFile != "" && keyFile != "" {
			log.Printf("binding to %s for secure HTTP server", *httpAddr)
			if err := serve.ListenAndServeTLS(certFile, keyFile); err != nil {
				log.Print(err)
			}
		} else {
			log.Printf("binding to %s for HTTP server", *httpAddr)
			if err := serve.ListenAndServe(); err != nil {
				log.Print(err)
			}
		}
	}()

	// Block/Wait for an error
	if err := <-errs; err != nil {
		shutdownServer()
		log.Print(err)
	}
}

func addPingRoute(r *mux.Router) {
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		moovhttp.SetAccessControlAllowHeaders(w, r.Header.Get("Origin"))
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PONG"))
	})
}

func addAssetsPath(r *mux.Router, assetPath string) {
	if _, err := os.Stat(assetPath); err != nil {
		panic(fmt.Sprintf("ERROR: unable to stat %s: %v", assetPath, err))
	}
	r.Methods("GET").PathPrefix("/").Handler(http.StripPrefix(*flagBasePath, http.FileServer(http.Dir(assetPath))))
}
