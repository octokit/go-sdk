#!/usr/bin/env bash

# GitHub Actions runs linting checks from .github/workflows/lint.yml.
# This script may be used to detect issues before pushing to a PR check.

# note: check the version of golangci-lint used in .github/workflows/lint.yml to ensure accuracy
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55

golangci-lint run ./...
