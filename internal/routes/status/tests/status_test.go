package status

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drajk/go-bsb-service/internal/config"
	"github.com/drajk/go-bsb-service/internal/routes/status"
	"github.com/stretchr/testify/assert"
)

func Test_StatusRoute_Should_Return_200(t *testing.T) {
	server := config.NewServer()
	router := server.WithRoutes("", status.Route(nil))
	req, _ := http.NewRequest(http.MethodGet, status.Path, nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
