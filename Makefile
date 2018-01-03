.PHONY: run build test

export GOPATH := ${PWD}

default: run

run:
	@go run ./src/main.go

build:
	@go build ./src/main.go

test:
	@go test -v ./src/chapter1/*

