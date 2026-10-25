# 8. Changes reach `main` only through pull requests

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

A merge to `main` deploys to production ([ADR 7](0007-digitalocean-droplet-blue-green.md)). Every change, including content, should be checked before it goes live.

## Decision

- Branch protection on `main` blocks direct pushes. Changes arrive only through pull requests with passing checks.
- `.github/workflows/ci.yml` runs on every pull request: `go tool templ generate`, build, and `go test ./...`. Nothing deploys from CI.
- `.github/workflows/deploy.yml` runs on merge to `main` and handles the build and deploy.
- `.github/dependabot.yml` opens pull requests for dependency updates, which go through the same checks.

## Consequences

- Nothing reaches production without passing CI.
- Even a one-line content fix needs a branch and a pull request.
