// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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

var (
	httpAddr  = flag.String("http.addr", bind.HTTP("ach"), "HTTP listen address")
	adminAddr = flag.String("admin.addr", bind.Admin("ach"), "Admin HTTP listen address")

	flagLogFormat = flag.String("log.format", "", "Format for log lines (Options: json, plain")

	logger log.Logger

	svc     server.Service
	handler http.Handler
)

func main() {
	flag.Parse()

	// Setup logging, default to stdout
	if v := os.Getenv("LOG_FORMAT"); v != "" {
		*flagLogFormat = v
	}
	if *flagLogFormat == "json" {
		logger = log.NewJSONLogger(os.Stdout)
	} else {
		logger = log.NewLogfmtLogger(os.Stdout)
	}
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

	// Check to see if our -http.addr flag has been overridden
	if v := os.Getenv("HTTP_BIND_ADDRESS"); v != "" {
		*httpAddr = v
	}

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

	// Check to see if our -admin.addr flag has been overridden
	if v := os.Getenv("HTTP_ADMIN_BIND_ADDRESS"); v != "" {
		*adminAddr = v
	}

	// Admin server (metrics and debugging)
	adminServer := admin.NewServer(*adminAddr)
	adminServer.AddVersionHandler(ach.Version) // Setup 'GET /version'
	go func() {
		logger.Log("admin", fmt.Sprintf("listening on %s", adminServer.BindAddr()))
		if err := adminServer.Listen(); err != nil {
			err = fmt.Errorf("problem starting admin http: %v", err)
			logger.Log("admin", err)
			errs <- err
		}
	}()
	defer adminServer.Shutdown()

	// Start main HTTP server
	go func() {
		if certFile, keyFile := os.Getenv("HTTPS_CERT_FILE"), os.Getenv("HTTPS_KEY_FILE"); certFile != "" && keyFile != "" {
			logger.Log("startup", fmt.Sprintf("binding to %s for secure HTTP server", *httpAddr))
			if err := serve.ListenAndServeTLS(certFile, keyFile); err != nil {
				errs <- err
				logger.Log("exit", err)
			}
		} else {
			logger.Log("startup", fmt.Sprintf("binding to %s for HTTP server", *httpAddr))
			if err := serve.ListenAndServe(); err != nil {
				errs <- err
				logger.Log("exit", err)
			}
		}
	}()

	if err := <-errs; err != nil {
		shutdownServer()
		logger.Log("exit", err)
	}
}
