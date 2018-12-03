// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package http implements a core suite of HTTP functions for use inside Moov. These packages are designed to
// be used in production to provide insight without an excessive performance tradeoff.
//
// This package implements several opininated response functions (See Problem, InternalError) and stateless CORS
// handling under our load balancing setup. They may not work for you.
//
// This package also implements a wrapper around http.ResponseWriter to log X-Request-ID, timing and the resulting status code.
package http
