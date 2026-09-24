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
	"bytes"
	"testing"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/moov-io/base/log"
	"github.com/stretchr/testify/require"
)

func TestWriteTimeoutFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    time.Duration
		wantLog string
	}{
		{name: "default", want: 30 * time.Second},
		{name: "seconds", value: "360s", want: 360 * time.Second, wantLog: "Using 6m0s as HTTP write timeout"},
		{name: "minutes", value: "5m", want: 5 * time.Minute, wantLog: "Using 5m0s as HTTP write timeout"},
		{name: "fractional", value: "1.5s", want: 1500 * time.Millisecond, wantLog: "Using 1.5s as HTTP write timeout"},
		{name: "disabled", value: "0", want: 0, wantLog: "Using 0s as HTTP write timeout"},
		{name: "invalid", value: "invalid", want: 30 * time.Second, wantLog: "Unable to parse HTTP_WRITE_TIMEOUT"},
		{name: "missing unit", value: "360", want: 30 * time.Second, wantLog: "Unable to parse HTTP_WRITE_TIMEOUT"},
		{name: "overflow", value: "999999999999999999999s", want: 30 * time.Second, wantLog: "Unable to parse HTTP_WRITE_TIMEOUT"},
		{name: "negative", value: "-1s", want: 30 * time.Second, wantLog: "HTTP_WRITE_TIMEOUT must not be negative"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("HTTP_WRITE_TIMEOUT", tc.value)
			var output bytes.Buffer
			logger := log.NewLogger(kitlog.NewLogfmtLogger(&output))

			require.Equal(t, tc.want, writeTimeoutFromEnv(logger))
			if tc.wantLog == "" {
				require.Empty(t, output.String())
			} else {
				require.Contains(t, output.String(), tc.wantLog)
			}
		})
	}
}
