package healthcontroller

import (
	"net/http"

	"github.com/mariusfa/gofl/health"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", health.HealthCheck)
}