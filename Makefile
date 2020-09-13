BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

BIN := $(GOPATH)/bin/$(NAME)
LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)"

all : generate compile run

compile :
	@go build $(LDFLAGS) -o $(BIN) main.go

generate :
	@go generate ./...

run :
	$(BIN)
