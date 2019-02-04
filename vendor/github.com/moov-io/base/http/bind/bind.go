// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package bind returns well known HTTP local bind addresses for Moov services.
// The package is intended for services to use for discovery during local development.
//
// This package also returns admin ports, useable with the github.com/moov-io/base/admin
// package.
package bind

import (
	"fmt"
	"strconv"
	"strings"
)

// serviceBinds is a map between a service name and its local bind address.
// The returned values will always be of the form ":XXXX" where XXXX is a
// valid port above 1024.
var serviceBinds = map[string]string{
	// Never change existing records, just add new records.
	"ach":     ":8080",
	"auth":    ":8081",
	"paygate": ":8082",
	"x9":      ":8083",
	"ofac":    ":8084",
	"gl":      ":8085",
	"fed":     ":8086",
}

// HTTP returns the local bind address for a Moov service.
func HTTP(serviceName string) string {
	v, ok := serviceBinds[strings.ToLower(serviceName)]
	if !ok {
		return ""
	}
	return v
}

// Admin returns the local bind address for a Moov service's admin server.
// This server typically serves metrics and debugging endpoints.
func Admin(serviceName string) string {
	http := HTTP(serviceName)
	if http == "" {
		return ""
	}
	http = strings.TrimPrefix(http, ":")
	n, err := strconv.Atoi(http)
	if err != nil {
		return ""
	}
	n += 1000 // 90XX
	n += 10   // 909X
	return fmt.Sprintf(":%d", n)
}
