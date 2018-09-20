package server

import (
	"net/http"
	"testing"
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
