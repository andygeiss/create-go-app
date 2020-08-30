# create-go-app

Create Go apps with with zero configuration by using a single command.

#### Table of Contents

- [Installation]()
- [Creating an App]()
- [Project Structure]()

## Installation

    go get -u github.com/andygeiss/create-go-app

## Creating Go Apps

Create a backend server app:

    create-go-app -type app -name <name>

Create a Go binary project:

    create-go-app -type bin -name <name>

Create a Go library/module:

    create-go-app -type lib -name <name>
