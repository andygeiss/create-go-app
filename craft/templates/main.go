package templates

// MainGo4App ...
var MainGo4App = `package main

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

// MainGo4Bin ...
var MainGo4Bin = `package main

import (
	"context"
	"fmt"

	"{{ .Path }}/internal/status"
	"{{ .Path }}/pkg/api"
)

var (
	build string = "no-build"
	version string = "no-version"
)

func main() {
	res, err := status.Service()(context.Background(), &api.StatusRequest{})
	fmt.Printf("Error is %v\n", err)
	fmt.Printf("StatusResponse is %v\n", res)
}`
