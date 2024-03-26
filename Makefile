COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_TIME=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash
BIN_NAME="adx-admin"
APP_MAIN_DIR=cmd

all: build

.PHONY: build
build:
	go build -ldflags "-X main.GitTag=$(GIT_TAG) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(COMMIT_HASH) -X main.GitAuthor=somebody"  -o ${BIN_NAME} ./cmd

generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...


.PHONY: wire
# wire
wire:
	cd $(APP_MAIN_DIR) && wire
