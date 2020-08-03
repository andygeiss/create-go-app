package templates

// ServerGo ...
var ServerGo = `// Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

package server

import (
	"log"
	"net/http"
	"os"
	"sync"

	"{{ .Path }}/pkg/api"
)

type metrics struct {
	ErrorCount     map[string]int ` + "`json:\"error_count\"`" + `
	RequestCount   map[string]int ` + "`json:\"request_count\"`" + `
	ResponseTimeMs map[string]int ` + "`json:\"response_time_ms\"`" + `
	mutex          sync.Mutex
}

// Server ...
type Server struct {
	logger        *log.Logger
	metrics       *metrics
	router        *http.ServeMux
	statusService api.StatusService
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// WithStatusService ...
func (s *Server) WithStatusService(service api.StatusService) {
	s.statusService = service
}

// NewServer ...
func NewServer() *Server {
	srv := &Server{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		metrics: &metrics{
			ErrorCount:     make(map[string]int),
			RequestCount:   make(map[string]int),
			ResponseTimeMs: make(map[string]int),
		},
		router: http.NewServeMux(),
	}
	srv.routes()
	return srv
}`