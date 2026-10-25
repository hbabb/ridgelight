package main

import (
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
