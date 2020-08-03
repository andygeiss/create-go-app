package templates

// MainGo ...
var MainGo = `package main

import (
	"log"
	"net/http"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/server"
)

var (
	build string = "no-build"
	version string = "no-version"
)

func main() {
	srv := server.NewServer()
	srv.WithStatusService(status.Service())
	log.Printf("####### {{ .Name }} %s (%s): Start listening ...\n", version, build)
	log.Fatal(http.ListenAndServe(":3000", srv))
}`
