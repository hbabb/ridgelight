// Command server runs the Ridgelight web server.
package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"

	"github.com/hbabb/ridgelight"
	"github.com/hbabb/ridgelight/views"
)

func main() {
	mux, err := newMux()
	if err != nil {
		log.Fatalf("setting up routes: %v", err)
	}

	// Timeouts keep slow or stalled clients from holding connections open forever.
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("listening on http://localhost%s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

// newMux registers every route on a new ServeMux.
// It is separated from main() so that it can be used in tests.
func newMux() (*http.ServeMux, error) {
	// Strip the "static" folder prefix from the embedded files, so
	// static/css/main.css is served at /static/css/main.css instead of
	// /static/static/css/main.css
	staticFS, err := fs.Sub(ridgelight.Static, "static")
	if err != nil {
		return nil, fmt.Errorf("loading embedded static files: %w", err)
	}

	mux := http.NewServeMux()

	// "{$}" makes this match only the exact path "/".
	// Without it, "GET /" would catch every path that has no other route.
	mux.Handle("GET /{$}", templ.Handler(views.Home()))
	mux.Handle("GET /projects", templ.Handler(views.Projects()))
	mux.Handle("GET /resume", templ.Handler(views.Resume()))

	// Health check for the blue/green deploy: 200 means the server is up.
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(staticFS)))

	return mux, nil
}
