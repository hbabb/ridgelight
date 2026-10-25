# 10. Local tooling: Just, go.mod tools, and golangci-lint

- **Status:** Accepted
- **Date:** 2026-10-25

## Context

The project needs repeatable commands for formatting, code generation, linting, vulnerability scanning, testing, and running the server, both locally and in CI. Tool versions should stay current without manual updates.

mise was considered. It can install tools and run tasks from one `mise.toml` file. Here, though, Go already installs every Go-based tool through `go.mod`, which leaves mise covering only golangci-lint. Dependabot also doesn't support `mise.toml`, so tools listed there wouldn't get automatic update pull requests.

## Decision

- **Go tools live in `go.mod`.** templ and govulncheck are listed with the `tool` directive and run with `go tool templ` and `go tool govulncheck`. Dependabot keeps them updated.
- **The Go toolchain comes from `go.mod`.** The `go` line sets the required version. An older local Go downloads it automatically.
- **golangci-lint is installed from the system package manager** (`pacman` on Arch). Its docs recommend against installing it through `go tool`. In CI, `golangci/golangci-lint-action` installs it.
- **Linting and Go formatting are configured in `.golangci.yml`:** the standard linters plus errorlint, gosec, nilerr, bodyclose, noctx, usestdlibvars, gocritic, revive, unconvert, unparam, modernize, misspell, and nolintlint, with gofumpt and goimports as formatters.
- **Just runs the tasks** defined in `justfile`. `just pre-commit` formats, generates, lints, scans, and tests before each commit.

## Consequences

- A fresh clone needs Go, golangci-lint, and Just installed. Everything else comes from `go.mod`.
- Dependabot opens update pull requests for templ, govulncheck, and the CI actions.
- golangci-lint updates come from system updates locally and from `version: latest` in CI, so the two can briefly differ.
- CI runs the same checks with direct commands rather than through Just.
