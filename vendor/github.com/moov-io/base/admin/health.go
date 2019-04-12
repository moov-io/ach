// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

var (
	errTimeout = errors.New("timeout exceeded")

	httpCheckTimeout = 10 * time.Second
)

type healthCheck struct {
	name  string
	check func() error
}

// Error executes the health check and will block until the result returns
func (hc *healthCheck) Error() error {
	if hc == nil || hc.check == nil {
		return nil
	}
	return hc.check()
}

type result struct {
	name string
	err  error
}

// AddLivenessCheck will register a new health check that is executed on every
// HTTP request of 'GET /live' against the admin server.
//
// Every check will timeout after 10s and return a timeout error.
//
// These checks are designed to be unhealthy only when the application has started but
// a dependency is unreachable or unhealthy.
func (s *Server) AddLivenessCheck(name string, f func() error) {
	s.liveChecks = append(s.liveChecks, &healthCheck{
		name:  name,
		check: f,
	})
}

func (s *Server) livenessHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results := processChecks(s.liveChecks)
		if len(results) == 0 {
			w.WriteHeader(http.StatusOK)
			return
		}
		status := http.StatusOK
		kv := make(map[string]string)
		for i := range results {
			if results[i].err != nil {
				status = http.StatusBadRequest
				kv[results[i].name] = results[i].err.Error()
			} else {
				kv[results[i].name] = "good"
			}
		}
		bs, _ := json.Marshal(kv)
		w.WriteHeader(status)
		w.Write(bs)
	}
}

// AddReadinessCheck will register a new health check that is executed on every
// HTTP request of 'GET /ready' against the admin server.
//
// Every check will timeout after 10s and return a timeout error.
//
// These checks are designed to be unhealthy while the application is starting.
func (s *Server) AddReadinessCheck(name string, f func() error) {
	s.readyChecks = append(s.readyChecks, &healthCheck{
		name:  name,
		check: f,
	})
}

func (s *Server) readinessHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		results := processChecks(s.readyChecks)
		if len(results) == 0 {
			w.WriteHeader(http.StatusOK)
			return
		}
		status := http.StatusOK
		kv := make(map[string]string)
		for i := range results {
			if results[i].err != nil {
				status = http.StatusBadRequest
				kv[results[i].name] = results[i].err.Error()
			} else {
				kv[results[i].name] = "good"
			}
		}
		bs, _ := json.Marshal(kv)
		w.WriteHeader(status)
		w.Write(bs)
	}
}

func processChecks(checks []*healthCheck) []result {
	var results []result
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(checks))

	for i := range checks {
		go func(check *healthCheck) {
			defer wg.Done()
			err := try(func() error { return check.Error() }, httpCheckTimeout)
			mu.Lock()
			results = append(results, result{
				name: check.name,
				err:  err,
			})
			mu.Unlock()
		}(checks[i])
	}

	wg.Wait()
	return results
}

// try will attempt to call f, but only for as long as t. If the function is still
// processing after t has elapsed then errTimeout will be returned.
func try(f func() error, t time.Duration) error {
	answer := make(chan error)
	go func() {
		answer <- f()
	}()
	select {
	case err := <-answer:
		return err
	case <-time.After(t):
		return errTimeout
	}
}
