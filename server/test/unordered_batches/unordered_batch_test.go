package unordered_batches

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/moov-io/ach/server/test"

	"github.com/stretchr/testify/require"
)

func TestUnorderedBatches(t *testing.T) {
	bs, err := os.ReadFile("Moov-batch-numbers-issue.json")
	require.NoError(t, err)

	server := test.NewServer()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", bytes.NewReader(bs))
	req.Header.Set("Content-Type", "application/json")
	server.Handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.Contains(t, w.Body.String(), "Batch numbers must be in ascending order")

	// Try with the ValidateOpt
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/files/create?unorderedBatchNumbers=true", bytes.NewReader(bs))
	req.Header.Set("Content-Type", "application/json")
	server.Handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var response struct {
		ID string `json:"id"`
	}
	json.NewDecoder(w.Body).Decode(&response)
	require.NotEmpty(t, response.ID)

	// Try POST /validate
	var buf bytes.Buffer
	buf.WriteString(`{"allowUnorderedBatchNumbers": true}`)
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/validate", response.ID), &buf)
	server.Handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	// Try POST /validate with ?unorderedBatchNumbers
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/validate?unorderedBatchNumbers=true", response.ID), nil)
	server.Handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	// Try POST /validate with ?allowUnorderedBatchNumbers
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/validate?allowUnorderedBatchNumbers=true", response.ID), nil)
	server.Handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)
}
