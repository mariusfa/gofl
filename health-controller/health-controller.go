package healthcontroller

import (
	"net/http"

	"github.com/mariusfa/gofl/v2/health"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", health.HealthCheck)
}
