# 7. DigitalOcean Droplet with blue/green deploys

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

A live Go server ([ADR 2](0002-live-go-server.md)) needs a host that keeps it running. Deploys should not take the site down, and a bad release should be easy to undo.

## Decision

- Host on a DigitalOcean Droplet.
- Build in GitHub Actions. Nothing is compiled on the Droplet.
- Run two copies of the site as containers, **blue** and **green**, each defined as a Podman Quadlet `.container` file in `deploy/quadlet/`. Only one receives traffic at a time.
- Caddy is the reverse proxy in front of them (`deploy/Caddyfile`) and points at the live container.
- A deploy starts the new version in the idle container, waits for a health check to pass, then switches Caddy to it and reloads Caddy.
- The previous container stays available, so rolling back means switching Caddy back.

How the build gets from GitHub Actions to the Droplet (container image or plain binary) is still an open decision (see the README).

## Consequences

- Deploys don't interrupt visitors, and rollback is one Caddy switch.
- The Droplet needs Podman, Caddy, and enough memory for two copies of the site.
- The server needs a health check endpoint for the deploy to check.
