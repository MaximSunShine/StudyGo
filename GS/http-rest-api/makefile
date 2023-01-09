.PHONY: build
build:
    go build -v ./cmd/apiserver

.Phony: test
test:
    go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := built