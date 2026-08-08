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

package server

import (
	"net/http"
	"os"
	"strings"
	"sync"
)

// CORSAllowedOriginsEnv is a comma-separated list of exact Origins permitted for
// credentialed CORS. Example: "https://moov.io,https://dashboard.moov.io".
// Localhost (http://localhost:<port>) remains allowed for development.
//
// Temporary ACH-local allowlist until moov-io/base#509 lands and is bumped here.
const CORSAllowedOriginsEnv = "MOOV_CORS_ALLOW_ORIGINS"

var (
	corsAllowlistOnce sync.Once
	corsAllowlist     map[string]struct{}
)

func loadCORSAllowlist() map[string]struct{} {
	corsAllowlistOnce.Do(func() {
		corsAllowlist = make(map[string]struct{})
		for _, part := range strings.Split(os.Getenv(CORSAllowedOriginsEnv), ",") {
			origin := strings.TrimSpace(part)
			if origin != "" {
				corsAllowlist[origin] = struct{}{}
			}
		}
	})
	return corsAllowlist
}

// resetCORSAllowlistForTest clears the cached allowlist. Intended for tests only.
func resetCORSAllowlistForTest() {
	corsAllowlistOnce = sync.Once{}
	corsAllowlist = nil
}

func originAllowedForCORS(origin string) bool {
	if origin == "" {
		return false
	}
	if strings.HasPrefix(origin, "http://localhost:") {
		return true
	}
	_, ok := loadCORSAllowlist()[origin]
	return ok
}

// setAccessControlAllowHeaders writes Access-Control-Allow-* headers only for
// allowlisted Origins. Do not call moov-io/base's reflector until it gains an
// allowlist (base#509).
func setAccessControlAllowHeaders(w http.ResponseWriter, origin string) {
	if !originAllowedForCORS(origin) {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Cookie,X-User-Id,X-Request-Id,Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
