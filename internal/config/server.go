package config

import (
	"fmt"
	"net/http"

	"github.com/drajk/go-bsb-service/pkg/logger"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server - Container with mux router
type Server struct {
	*mux.Router
}

// ServerConfigOption - Struct to make server configurable
type ServerConfigOption func(server *Server)

//NewServer creates a new server
func NewServer(options ...ServerConfigOption) *Server {
	s := &Server{
		mux.NewRouter().StrictSlash(true),
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

//Start starts the server with given configuration
func (s *Server) Start(cfg ApplicationConfig) {
	// TODO: add tracers here
	// if cfg.TracingEnabled() {
	//     do something
	// }

	s.start("", cfg.ServerPort())
}

func (s *Server) start(addr string, port int) {
	logger.GetLogger().Info("starting server")

	sv := &http.Server{
		Addr: fmt.Sprintf("%s:%v", addr, port),
		Handler: handlers.RecoveryHandler(
			handlers.PrintRecoveryStack(true),
			handlers.RecoveryLogger(logger.GetLogger()))(s),
	}

	panic(sv.ListenAndServe())
}
