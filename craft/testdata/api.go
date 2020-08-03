package api

import "context"

//go:generate gocraft -type api

// StatusRequest ...
type StatusRequest struct {
}

// StatusResponse ...
type StatusResponse struct {
	Text string `json:"text"`
}

// StatusService ...
type StatusService func(ctx context.Context, req *StatusRequest) (res *StatusResponse, err error)
