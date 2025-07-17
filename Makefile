# Makefile for merriam-webster Go project

.PHONY: all build test lint tidy vendor install

all: build

build:
	go build -o bin/merriam-webster ./source/cmd/app

test:
	go test ./...

lint:
	./bin/golangci-lint run ./source/...

tidy:
	go mod tidy

vendor:
	go mod vendor 

GO_PACKAGE=github.com/example/merriam-webster

.PHONY: install
install:
	GOPATH=$(shell go env GOPATH) go install -mod vendor -ldflags="-X main.version=$(TAG)" $(GO_PACKAGE)/source/cmd/app 