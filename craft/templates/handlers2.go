package templates

// Handlers2Go ...
var Handlers2Go = `// Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

package server

import (
	"context"
	"encoding/json"
	"net/http"

	"{{ .Path }}/pkg/api"
)

{{ range $i, $name := .Services }}
func (s *Server) handle{{ $name }}() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request from the client.
		req := new(api.{{ $name }}Request)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.increaseErrorCount(r)
			s.logger.Printf("%-6s %-20s %-50s %-s", "ERROR", r.RemoteAddr, r.RequestURI, err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Publish the event.
		s.bus.Publish("{{ lc $name }}", req)
		// Call the service.
		res, err := s.{{ lc $name }}Service(context.Background(), req)
		if err != nil {
			s.increaseErrorCount(r)
			s.logger.Printf("%-6s %-20s %-50s %-s", "ERROR", r.RemoteAddr, r.RequestURI, err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// Publish the event.
		s.bus.Publish("{{ lc $name }} done", req)
		// Encode the response and send it to the client.
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			s.increaseErrorCount(r)
			s.logger.Printf("%-6s %-20s %-50s %-s", "ERROR", r.RemoteAddr, r.RequestURI, err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
{{ end }}

func (s *Server) handleMetrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Encode the response and send it to the client.
		if err := json.NewEncoder(w).Encode(s.metrics); err != nil {
			s.logger.Printf("%-6s %-20s %-50s %-s", "ERROR", r.RemoteAddr, r.RequestURI, err.Error())
		}
	}
}`
