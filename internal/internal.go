package internal

import (
	"github.com/drajk/go-bsb-service/internal/config"
	"github.com/drajk/go-bsb-service/internal/routes/status"
)

// SetupServer - Bootstrap configures and runs the application
func SetupServer(appContainer config.IContainer) *config.Server {
	server := config.NewServer().
		WithRoutes("", status.Route(appContainer))

	return server
}
