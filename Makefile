NAME             := merriam-webster
NAMESPACE        := merriam-webster
DOCKER_IMAGE     := $(NAMESPACE)/$(NAME)
GOLANGCI_VERSION := v1.21.0
GO_PACKAGE=github.com/example/merriam-webster
MOCKERY_VERSION  := v2
TAG              := $(shell git rev-parse --short HEAD)
.DEFAULT_GOAL    := install
.PHONY: all build test lint tidy vendor install

build:
	$(info --- Building local binary)
	go build -o bin/merriam-webster ./source/cmd/merriam-webster

test:
	$(info --- Running cli tool tests)
	go test -v ./...

lint:
	$(info --- Running Go lint checks)
	./bin/golangci-lint run ./source/...

tidy:
	go mod tidy

vendor:
	go mod vendor 

.PHONY: push
push: build
ifeq ($(CI),true)
	$(info --- Pushing image to the Artifactory registry)
	#   connect to artifactory code
else
	$(warning This is only allowed to be run in CI)
endif	

.PHONY: install
install: build
	$(info --- Installing tool to your system)
	sudo cp bin/merriam-webster /usr/local/bin/merriam-webster
	@echo "Installed merriam-webster to to your system"
	@echo "Try running 'merriam-webster --help' to get started"
