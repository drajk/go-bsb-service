package status

import (
	"net/http"

	"github.com/drajk/go-bsb-service/internal/config"
)

type container interface {
}

// Path - status path
const Path = "/status"

// Route Function - /status healthcheck endpoint
func Route(appContainer container) (route config.Route) {
	route = config.Route{
		Path:    Path,
		Method:  http.MethodGet,
		Handler: handler(),
	}
	return route
}
