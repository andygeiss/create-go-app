# create-go-app

Create Go apps with with zero configuration by using a single command.

#### Table of Contents

- [Installation]()
- [Creating an App]()
- [Project Structure]()

## Installation

    go get -u github.com/andygeiss/create-go-app

## Creating an App

    create-go-app -type app -name <app name>

## Project Structure

    ├── go.mod
    ├── internal
    │   └── status
    │       ├── service.go
    │       └── service_test.go
    ├── main.go
    └── pkg
        ├── api
        │   ├── api.go
        │   └── api.http
        ├── assert
        │   └── assert.go
        └── server
            ├── handlers.go
            ├── middleware.go
            ├── routes.go
            └── server.go
