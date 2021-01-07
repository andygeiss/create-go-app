package templates

// MainGo4App ...
var MainGo4App = `package main

import (
	"log"
	"net/http"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/event"
	"{{ .Path }}/pkg/server"
)

var (
	build string = "no-build"
	version string = "no-version"
)

func main() {
	bus := event.NewBus()
	srv := server.NewServer(bus)
	srv.WithStatusService(status.Service(repository.NewFileRepository("repository.json")))
	log.Printf("####### {{ .Name }} %s (%s): Start listening ...\n", version, build)
	log.Fatal(http.ListenAndServe(":3000", srv))
}`

// MainGo4Bin ...
var MainGo4Bin = `package main

import (
	"context"
	"fmt"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/api"
)

var (
	build 	string = "no-build"
	name    string = "no-name"
	version string = "no-version"
)

func main() {
	res, err := status.Service()(context.Background(), &api.StatusRequest{})
	fmt.Printf("%s %s (%s)\n", name, version, build)
	fmt.Printf("error is %v\n", err)
	fmt.Printf("response is %v\n", res)
}`
