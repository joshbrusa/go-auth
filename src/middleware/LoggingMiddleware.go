package middleware

import (
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(
			"request",
			"path", r.URL.Path,
			"method", r.Method,
		)
		next.ServeHTTP(w, r)
	})
}
