# AGENTS.md

## Project Overview

`tz` is a lightweight, zero-dependency Go package that maps IANA timezone identifiers to ISO 3166-1 alpha-2 country codes and UTC offsets. Module path: `github.com/infobits-io/tz`.

## Project Structure

```
tz.go          # Package API: Decode, IsValid, All, ByCountryCode, ByUtcOffset, Current
data.go        # Timezone data map (400+ IANA entries)
tz_test.go     # Tests, benchmarks, and examples
go.mod         # Go 1.26, no external dependencies
Makefile       # Development commands (lint, test, build, coverage, bench)
.golangci.yml  # Linter configuration (golangci-lint v2)
```

## Development Commands

```bash
make lint      # Run golangci-lint (includes auto-fix)
make test      # Run tests with -v -race
make build     # Build the package
make coverage  # Run tests with coverage report
make bench     # Run benchmarks with memory stats
make clean     # Remove generated artifacts
make help      # Print available targets
```

## Code Conventions

- **Go version**: 1.26
- **Formatting**: Enforced by golangci-lint via gofmt, gofumpt, goimports, golines (max 200 chars)
- **Linting**: Strict golangci-lint v2 config with many linters enabled; line length limit is 200
- **Testing**: Use `-race` flag; test files exclude errcheck, dupl, gosec linters
- **No external dependencies**: Keep this package dependency-free
- **Comments**: End with a period (enforced by `godot` linter)

## Workflow

1. Run `make lint` before committing
2. Run `make test` to verify changes
3. CI runs lint, test, and build as separate jobs (Go 1.26)
