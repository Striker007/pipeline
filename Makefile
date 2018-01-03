.PHONY: run test

export GOPATH := ${PWD}

default: run

run:
	@go run ./src/main.go

test:
	@go test -v ./src/chapter1/*

