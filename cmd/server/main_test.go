package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoutes(t *testing.T) {
	mux, err := newMux()
	require.NoError(t, err)

	tests := []struct {
		path string
		want int
	}{
		{"/", http.StatusOK},
		{"/projects", http.StatusOK},
		{"/resume", http.StatusOK},
		{"/healthz", http.StatusOK},
		{"/static/css/main.css", http.StatusOK},
		{"/does-not-exist", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			assert.Equal(t, tt.want, rec.Code, "GET %s", tt.path)
		})
	}
}

// TestServe starts the real server the same way main does
// and checks that each page answers over a real network connection.
func TestServe(t *testing.T) {
	var lc net.ListenConfig
	ln, err := lc.Listen(t.Context(), "tcp", "127.0.0.1:0")
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan error, 1)
	go func() {
		done <- serve(ctx, ln)
	}()

	base := "http://" + ln.Addr().String()
	for _, path := range []string{"/", "/projects", "/resume", "/healthz"} {
		req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, base+path, nil)
		require.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())

		assert.Equal(t, http.StatusOK, resp.StatusCode, "GET %s", path)
	}

	cancel()
	require.NoError(t, <-done, "server should shut down cleanly")
}
