// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"strings"
	"testing"
)

func TestFiles__createBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory()
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := createBatchEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(createBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}

func TestFiles__getBatchesEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory()
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getBatchesEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(getBatchesResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}

func TestFiles__getBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory()
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getBatchEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(getBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}

func TestFiles__deleteBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory()
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := deleteBatchEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(deleteBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}
