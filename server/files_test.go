// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"strings"
	"testing"
)

func TestFiles__createFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := createFileEndpoint(svc, repo, nil)(context.TODO(), body)
	r, ok := resp.(createFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

func TestFiles__getFilesEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFilesEndpoint(svc)(context.TODO(), body)
	_, ok := resp.(getFilesResponse)
	if !ok || err != nil {
		t.Errorf("got %#v : err=%v", resp, err)
	}
}

func TestFiles__getFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFileEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(getFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

func TestFiles__getFileContentsEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFileContentsEndpoint(svc, nil)(context.TODO(), body)
	_, ok := resp.(getFileContentsResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil {
		t.Errorf("expected error: err=%v", err)
	}

}

func TestFiles__validateFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := validateFileEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(validateFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}
