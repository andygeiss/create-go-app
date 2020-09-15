package templates

// APIGo ...
var APIGo = `package api

import "context"

//go:generate {{ .Generator }} -type api -name {{ .Name }} -path {{ .Path }}
//go:generate {{ .Generator }} -type bundle -name {{ .Name }} -path {{ .Path }}

// StatusRequest ...
type StatusRequest struct {
}

// StatusResponse ...
type StatusResponse struct {
	Text string ` + "`json:\"text\"`" + `
}

// StatusService ...
type StatusService func(ctx context.Context, req *StatusRequest) (res *StatusResponse, err error)`
