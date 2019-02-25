// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/server"
	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/http/bind"

	"github.com/go-kit/kit/log"
)

/**
CreateFile
curl -d '{"id":"1234","immediateDestination":"9876543210", "immediateDestinationName":"Federal Reserve Bank", "immediateOrigin":"1234567890","immediateOriginName":"My Bank Name"}' -H "Content-Type: application/json" -X POST http://localhost:8080/files/create

GetFile
curl -H "Content-Type: application/json" -X GET http://localhost:8080/files/1234

GetFiles
curl -H "Content-Type: application/json" -X GET http://localhost:8080/files/

DeleteFile
curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/files/1234

CreateBatch
curl -d '{"id":"54321","serviceClassCode":220,"standardEntryClassCode":"WEB","companyName":"Your Company inc","companyIdentification":"121042882","companyEntryDescription":"Online Order","ODFIIdentification":"12104288"}' -H "Content-Type: application/json" -X POST http://localhost:8080/files/08B751B2/batches/

GetBatch
curl -H "Content-Type: application/json" -X GET http://localhost:8080/files/08B751B2/batches/54321

GetBatches
curl -H "Content-Type: application/json" -X GET http://localhost:8080/files/08B751B2/batches/

DeleteBatch
curl -H "Content-Type: application/json" -X DELETE http://localhost:8080/files/08B751B2/batches/54321
**/

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ach"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ach"), "Admin HTTP listen address")

	logger log.Logger

	svc     server.Service
	handler http.Handler
)

func main() {
	flag.Parse()

	// Setup logging, default to stdout
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger.Log("startup", fmt.Sprintf("Starting ach server version %s", ach.Version))

	// Setup underlying ach service
	var achFileTTL time.Duration
	if v := os.Getenv("ACH_FILE_TTL"); v != "" {
		dur, err := time.ParseDuration(v)
		if err == nil {
			achFileTTL = dur
			logger.Log("main", fmt.Sprintf("Using %v as ach.File TTL", achFileTTL))
		}
	}
	r := server.NewRepositoryInMemory(achFileTTL, logger)
	svc = server.NewService(r)

	// Create HTTP server
	handler = server.MakeHTTPHandler(svc, r, log.With(logger, "component", "HTTP"))

	// Listen for application termination.
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	readTimeout, _ := time.ParseDuration("30s")
	writTimeout, _ := time.ParseDuration("30s")
	idleTimeout, _ := time.ParseDuration("60s")

	serve := &http.Server{
		Addr:    *httpAddr,
		Handler: handler,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:  readTimeout,
		WriteTimeout: writTimeout,
		IdleTimeout:  idleTimeout,
	}
	shutdownServer := func() {
		if err := serve.Shutdown(context.TODO()); err != nil {
			logger.Log("shutdown", err)
		}
	}

	// Admin server (metrics and debugging)
	adminServer := admin.NewServer(*adminAddr)
	go func() {
		logger.Log("admin", fmt.Sprintf("listening on %s", adminServer.BindAddr()))
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			logger.Log("admin", err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- serve.ListenAndServe()
		// TODO(adam): support TLS
		// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
	}()

	if err := <-errs; err != nil {
		shutdownServer()
		logger.Log("exit", err)
	}
}
