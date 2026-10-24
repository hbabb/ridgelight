# 3. templ for HTML templates

- **Status:** Accepted
- **Date:** 2026-10-24

## Context

Pages need shared layout and reusable pieces. Go's built-in `html/template` package checks templates only when they run, so a typo in a field name shows up as a runtime error.

## Decision

Use [templ](https://templ.guide) for all HTML.

- Components live in `views/` as `.templ` files.
- `go tool templ generate` compiles each `.templ` file into a `*_templ.go` file. These are ordinary Go functions, so the Go compiler checks them.
- templ is tracked as a project tool with the `tool` directive in `go.mod` (Go 1.24+). The version is pinned per project, and Dependabot can update it.

## Consequences

- Template mistakes are caught at compile time.
- There is a code generation step before building. How the generated files are handled is covered in [ADR 9](0009-commit-generated-templ-files.md).
- The templ CLI (from `go tool`) and the templ runtime library (imported by the generated code) both come from the same module version in `go.mod`.
