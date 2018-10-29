package server

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/moov-io/ach"

	httptransport "github.com/go-kit/kit/transport/http"
)

func TestEncodeResponse(t *testing.T) {
	ctx := context.TODO()
	w := httptest.NewRecorder()
	if err := encodeResponse(ctx, w, "hi mom"); err != nil {
		t.Fatal(err)
	}
	w.Flush()

	var resp string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	if resp != "hi mom" {
		t.Errorf("got %q", resp)
	}

	v := w.Header().Get("content-type")
	if v != "application/json; charset=utf-8" {
		t.Errorf("got %q", v)
	}
}

func TestEncodeTextResponse(t *testing.T) {
	ctx := context.TODO()
	w := httptest.NewRecorder()
	if err := encodeTextResponse(ctx, w, strings.NewReader("hi mom")); err != nil {
		t.Fatal(err)
	}
	if v := w.Body.String(); v != "hi mom" {
		t.Errorf("got %q", v)
	}

	if v := w.Header().Get("content-type"); v != "text/plain" {
		t.Errorf("got %q", v)
	}
}

func TestXTotalCountHeader(t *testing.T) {
	counter := getFilesResponse{
		Files: []*ach.File{ach.NewFile()},
		Err:   nil,
	}

	w := httptest.NewRecorder()
	encodeResponse(context.Background(), w, counter)

	actual, ok := w.Result().Header["X-Total-Count"]
	if !ok {
		t.Fatal("should have count")
	}
	if actual[0] != "1" {
		t.Errorf("should be 1, got %v", actual[0])
	}
}

func TestRouting__proxyCORSHeaders(t *testing.T) {
	r := httptest.NewRequest("GET", "/ping", nil)
	r.Header.Set("Access-Control-Allow-Origin", "origin")
	r.Header.Set("Access-Control-Allow-Methods", "methods")
	r.Header.Set("Access-Control-Allow-Headers", "headers")
	r.Header.Set("Access-Control-Allow-Credentials", "credentials")

	ctx := context.TODO()
	ctx = saveCORSHeadersIntoContext()(ctx, r)

	check := func(ctx context.Context, key contextKey, expected string) {
		v, ok := ctx.Value(key).(string)
		if !ok {
			t.Errorf("key=%v, v=%s, ok=%v", key, v, ok)
		}
		if v != expected {
			t.Errorf("got %s, expected %s", v, expected)
		}
	}

	check(ctx, accessControlAllowOrigin, "origin")
	check(ctx, accessControlAllowMethods, "methods")
	check(ctx, accessControlAllowHeaders, "headers")
	check(ctx, accessControlAllowCredentials, "credentials")

	// now make sure ctx writes these headers to an http.ResponseWriter
	w := httptest.NewRecorder()
	respondWithSavedCORSHeaders()(ctx, w)
	w.Flush()

	if v := r.Header.Get("Access-Control-Allow-Origin"); v != "origin" {
		t.Errorf("got %s", v)
	}
	if v := r.Header.Get("Access-Control-Allow-Methods"); v != "methods" {
		t.Errorf("got %s", v)
	}
	if v := r.Header.Get("Access-Control-Allow-Headers"); v != "headers" {
		t.Errorf("got %s", v)
	}
	if v := r.Header.Get("Access-Control-Allow-Credentials"); v != "credentials" {
		t.Errorf("got %s", v)
	}
}

func TestPreflightHandler(t *testing.T) {
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(saveCORSHeadersIntoContext()),
		httptransport.ServerAfter(respondWithSavedCORSHeaders()),
	}

	handler := preflightHandler(options)

	// Make our pre-flight request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("OPTIONS", "/files/create", nil)

	// Add CORS headers (i.e. like they're coming from auth)
	r.Header.Set("Access-Control-Allow-Origin", "origin")
	r.Header.Set("Access-Control-Allow-Methods", "methods")
	r.Header.Set("Access-Control-Allow-Headers", "headers")
	r.Header.Set("Access-Control-Allow-Credentials", "credentials")

	// Make the request
	handler.ServeHTTP(w, r)
	w.Flush()

	// Check response
	if v := w.Header().Get("Access-Control-Allow-Origin"); v != "origin" {
		t.Errorf("got %s", v)
	}
	if v := w.Header().Get("Access-Control-Allow-Methods"); v != "methods" {
		t.Errorf("got %s", v)
	}
	if v := w.Header().Get("Access-Control-Allow-Headers"); v != "headers" {
		t.Errorf("got %s", v)
	}
	if v := w.Header().Get("Access-Control-Allow-Credentials"); v != "credentials" {
		t.Errorf("got %s", v)
	}
}
