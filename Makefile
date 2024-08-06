BINARY := yatter-backend-go
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PATH := $(PATH):${MAKEFILE_DIR}bin
SHELL := env PATH="$(PATH)" /bin/bash

GOARCH = amd64

build:
	CGO_ENABLED=0 go build -o build/${BINARY}

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/${BINARY}-linux-${GOARCH} .

mod:
	go mod download

test:
	go test $(shell go list ${MAKEFILE_DIR}/...)

#https://github.com/golangci/golangci-lint/issues/2649
lint:
	if ! [ -x $(GOPATH)/bin/golangci-lint ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2; \
	fi
	golangci-lint run --concurrency 2

vet:
	go vet ./...

.PHONY:	build mod test lint vet clean
