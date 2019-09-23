SHELL := /bin/bash

default: lint build test

.PHONY: build
build:
	go build ./...

.PHONY: ci
ci: install.deps lint build test

.PHONY: install.deps
install.deps: install.golangci-lint

.PHONY: install.golangci-lint
install.golangci-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $$(go env GOPATH)/bin latest

.PHONY: lint
lint:
	golangci-lint run ./... --enable-all --disable gochecknoglobals,lll --deadline 10m

.PHONY: test
test:
	go test -count 1 -v ./...
