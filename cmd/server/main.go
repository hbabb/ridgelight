// Command server runs the Ridgelight web server.
package main

import (
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"

	"github.com/hbabb/ridgelight"
	"github.com/hbabb/ridgelight/views"
)

func main() {
	// Strip the "static" folder prefix from the embedded files, so
	// static/css/main.css is served at /static/css/main.css instead of
	// /static/static/css/main.css.
	staticFS, err := fs.Sub(ridgelight.Static, "static")
	if err != nil {
		log.Fatalf("loading embedded static files: %v", err)
	}

	mux := http.NewServeMux()

	// "{$}" makes this match only the exact path "/".
	// Without it, "GET /" would catch every path that has no other route.
	mux.Handle("GET /{$}", templ.Handler(views.Home()))

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(staticFS)))

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
