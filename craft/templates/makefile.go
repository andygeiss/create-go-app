package templates

// Makefile ...
var Makefile = `

BUILD := $(shell git rev-parse --short HEAD)
MODULE := $(shell cut -f2 -d" " go.mod | head -1)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

BIN := $(GOPATH)/bin/$(NAME)
LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)"

all : compile run

compile :
	@go build $(LDFLAGS) -o $(BIN) main.go
	
licenses :
	@go-licenses csv $(MODULE) > LICENSE.csv

setup :
	@go get -u github.com/google/go-licenses
	@go install github.com/google/go-licenses

run :
	$(BIN)

`
