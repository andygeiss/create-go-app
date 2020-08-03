package templates

// Server2Go ...
var Server2Go = `// Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

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
	{{ range $i, $name := .Services }}{{ lc $name }}Service api.{{ $name }}Service
	{{ end }}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

{{ range $i, $name := .Services }}
// With{{ $name }}Service ...
func (s *Server) With{{ $name }}Service(service api.{{ $name }}Service) {
	s.{{ lc $name }}Service = service
}
{{ end }}

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