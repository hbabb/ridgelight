# List all tasks
default:
    @just --list

# Format .templ files
templ-fmt:
    go tool templ fmt .

# Generate Go code from .templ files
generate:
    go tool templ generate

# Format Go files (gofumpt and goimports)
fmt:
    golangci-lint fmt

# Run the linters
lint:
    golangci-lint run

# Check for known vulnerabilities
vuln:
    go tool govulncheck ./...

# Run the tests
test:
    go test ./...

# Run the server at http://localhost:8080
run:
    go run ./cmd/server

# Format, generate, lint, scan, and test before commiting
pre-commit: templ-fmt generate fmt lint vuln test