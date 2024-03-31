package accessmiddleware

import (
	"net/http"
	"time"

	accesslog "github.com/mariusfa/gofl/v2/logging/access-log"
)

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *customResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	// Defaults to 200 status code
	return &customResponseWriter{w, http.StatusOK}
}

// This assumes access logger is already initialized.
func AccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		crw := newCustomResponseWriter(w)
		requestMethod := r.Method
		requestPath := r.URL.Path
		start := time.Now()
		next.ServeHTTP(crw, r)
		durationMs := int(time.Since(start) / time.Millisecond)

		accesslog.AccessLog.Info(durationMs, crw.status, requestPath, requestMethod)
	})
}
