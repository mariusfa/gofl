package metricscontroller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMetricsController(t *testing.T) {
	m := NewMetricsController()

	router := http.NewServeMux()

	m.RegisterRoutes(router)

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
