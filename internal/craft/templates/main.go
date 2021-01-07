package templates

// MainGo4Bin ...
var MainGo4Bin = `package main

import (
	"log"
	"net/http"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/event"
	"{{ .Path }}/pkg/repository"
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
