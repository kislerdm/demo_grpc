SHELL := /bin/bash

DIR_PROTO_BE := $(PWD)/backend/proto
DIR_PROTO_FE := $(PWD)/frontend/client/proto

help: ## Prints help message.
	@ grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1m%-30s\033[0m %s\n", $$1, $$2}'

proto-backend: ## Compile protobuf for the backend.
	@ test -d $(DIR_PROTO_BE) || mkdir $(DIR_PROTO_BE)
	@ protoc --go_out=$(DIR_PROTO_BE) --go_opt=paths=source_relative \
		--go-grpc_out=$(DIR_PROTO_BE) --go-grpc_opt=paths=source_relative \
		./app.proto

proto-frontend: ## Compile protobuf for the frontend.
	@ test -d $(DIR_PROTO_FE) || mkdir $(DIR_PROTO_FE)
	@ protoc --js_out=$(DIR_PROTO_FE) --js_opt=paths=source_relative \
		--js-grpc_out=$(DIR_PROTO_FE) --js-grpc_opt=paths=source_relative \
		./app.proto

build-backend: proto-backend ## builds the binary for the local arch of the backend.
	@ cd $(PWD)/backend \
		&& test -d bin || mkdir bin \
		&& CGO_ENABLED=0 GOARCH=$(uname -m) GOOS=$(uname) go build -a -gcflags=all="-l -B -C" -ldflags="-w -s" -o ./bin/runner *.go

test-backend: ## Runs tests for the backend.
	@ go test -v ./backend/...

run-backend: ## Launches the compiled binary for the backend.
	@ ./backend/bin/runner

compress: ## Compresses binary with upx.
	@ upx -9 --brute ./backend/bin/runner

.DEFAULT_GOAL := help
