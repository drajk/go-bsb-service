package config

import (
	"fmt"
	"net/http"

	"github.com/drajk/go-bsb-service/pkg/logger"
)

// Route - Struct used by WithRoutes
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// WithRoutes includes the routes as part of the server
func (s *Server) WithRoutes(basePath string, routes ...Route) *Server {
	sub := s.PathPrefix(basePath).Subrouter()

	for _, route := range routes {
		sub.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		logger.With(
			route.Method, fmt.Sprintf("%s%s", basePath, route.Path),
		).Infof("registered path")
	}
	return s
}
