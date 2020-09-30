BUILD := $(shell git rev-parse --short HEAD)
NAME := $(shell basename "$(PWD)")
VERSION := $(shell git describe --tags)

BIN := $(GOPATH)/bin/$(NAME)
LDFLAGS=-ldflags "-s -w -X=main.build=$(BUILD) -X=main.name=$(NAME) -X=main.version=$(VERSION)"

all : compile run

compile :
	@go build $(LDFLAGS) -o $(BIN) main.go

run :
	$(BIN)

setup :
	@go get -u github.com/google/go-licenses
	@go install github.com/google/go-licenses
	@sudo apt-get install -y git default-jre npm sassc
	@sudo npm -g install autoprefixer npm postcss-cli 
	@cp tools/google/closure-compiler/*.* $(GOPATH)/bin/
