// Command server runs the Ridgelight web server.
package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/a-h/templ"

	"github.com/hbabb/ridgelight"
	"github.com/hbabb/ridgelight/views"
)

// addr is the address the server listens on.
const addr = ":8080"

func main() {
	// ctx is canceled when the process receives Ctrl+C (SIGINT) or SIGTERM,
	// which is what Podman sends when it stops a container.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := run(ctx)
	stop()

	if err != nil {
		log.Fatal(err)
	}
}

// run opens the listener on addr and serves until ctx is canceled.
func run(ctx context.Context) error {
	var lc net.ListenConfig
	ln, err := lc.Listen(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("listening on %s: %w", addr, err)
	}

	return serve(ctx, ln)
}

// serve runs the HTTP server on ln until ctx is canceled,
// then shuts down gracefully, letting in-flight requests finish.
// It is separated from run so tests can pass their own listener.
func serve(ctx context.Context, ln net.Listener) error {
	mux, err := newMux()
	if err != nil {
		return fmt.Errorf("setting up routes: %w", err)
	}

	// Timeouts keep slow or stalled clients from holding connections open forever.
	srv := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	serveErr := make(chan error, 1)
	go func() {
		log.Printf("listening on %s", ln.Addr())
		serveErr <- srv.Serve(ln)
	}()

	select {
	case err := <-serveErr:
		return fmt.Errorf("serving: %w", err)
	case <-ctx.Done():
	}

	// ctx is already canceled, so the shutdown gets its own 10-second deadline.
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutting down: %w", err)
	}

	if err := <-serveErr; !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("serving: %w", err)
	}

	log.Print("server stopped")
	return nil
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
