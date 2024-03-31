package metricscontroller

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsController struct{}

func NewMetricsController() *MetricsController {
	return &MetricsController{}
}

func (m *MetricsController) RegisterRoutes(router *http.ServeMux) {
	router.Handle("/metrics", promhttp.Handler())
}
