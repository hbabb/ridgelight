# 9. Commit generated templ files

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

templ compiles each `.templ` file into a `*_templ.go` file ([ADR 3](0003-templ-for-html.md)). Those generated files can either be committed to the repo or ignored and regenerated before every build.

## Decision

Commit the generated `*_templ.go` files alongside their `.templ` sources.

## Consequences

- A fresh clone builds with `go build` and runs with `go run ./cmd/server`, without running templ first.
- After editing any `.templ` file, run `go tool templ generate` before committing. Otherwise the committed Go code is out of date.
- Pull requests show changes to both the `.templ` file and its generated file.
- CI runs `go tool templ generate` ([ADR 8](0008-pull-requests-only-to-main.md)). If that changes any files, the committed generated code was stale.
