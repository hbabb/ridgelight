# 2. Live Go server, not a static site

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

The site needs a projects page, a blog, and a resume page. Go is the preferred language for this project and is part of the author's learning path. A static site generator was considered and ruled out.

## Decision

Ridgelight is a live Go web server built with the standard library's `net/http` package. It compiles to a single binary that serves every page.

- Routing uses the standard library's `http.ServeMux` with method and path patterns (for example `GET /{$}`). No third-party router.
- The `http.Server` sets read, write, header, and idle timeouts so slow or stalled clients can't hold connections open indefinitely.
- The server listens on port `8080`.

## Consequences

- The site needs a host that keeps a process running. See [ADR 7](0007-digitalocean-droplet-blue-green.md).
- No dependencies are needed for routing or serving.
- Because content is baked into the binary ([ADR 5](0005-content-in-repo-no-database.md)), a content change means a new build and deploy.
