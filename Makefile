SHELL := /bin/bash

DIR_BIN := $(PWD)/bin
APP := server

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

proto: ## Compile protobuf.
	@ protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./app.proto

build: ## builds the binary for the local arch. Set APP variable to "server", or "client" to build the right app (default: server).
	@ test -d $(DIR_BIN) || mkdir $(DIR_BIN)
	@ CGO_ENABLED=0 GOARCH=$(uname -m) GOOS=$(uname) go build -a -gcflags=all="-l -B -C" -ldflags="-w -s" -o $(DIR_BIN)/$(APP) ./cmd/$(APP)/*.go

test: ## Runs tests.
	@ go test -v ./...

run: ## Launches the compiled binary. Set APP variable to "server", or "client" to build the right app (default: server).
	@ $(DIR_BIN)/$(APP)

compress: ## Compresses binary with upx. Set APP variable to "server", or "client" to build the right app (default: server).
	@ upx -9 --brute $(DIR_BIN)/$(APP)

.DEFAULT_GOAL := help
