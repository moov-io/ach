package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/ach"
)

func TestAcceptableContentLength(t *testing.T) {
	h := make(http.Header)

	if acceptableContentLength(h) { // reject if missing header
		t.Error("wanted unacceptable")
	}

	h.Set("Content-Length", "1000")
	if !acceptableContentLength(h) {
		t.Error("should have accepted")
	}

	h.Set("Content-Length", "10000000000000")
	if acceptableContentLength(h) {
		t.Error("should have rejected")
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
