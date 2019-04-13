package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Status OK")
	}))
	defer ts.Close()
	res, err := Get(ts.URL, nil)
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "Status OK", string(res))
}

func TestGetResponseWithHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "value", r.Header.Get("key"), "Expected value in header key")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "Status OK")
	}))
	defer ts.Close()
	header := map[string]string{"key": "value"}
	res, err := Get(ts.URL, header)
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "Status OK", string(res))
}
