package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"log/slog"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: 200}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func Logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger = logger.With(
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"clientId", r.Header.Get("X-Client-Id"),
				"requestId", r.Header.Get("X-Request-Id"),
			)

			start := time.Now()

			defer func() {
				if recovered := recover(); recovered != nil {
					err, ok := recovered.(error)

					w.WriteHeader(http.StatusInternalServerError)

					if ok {
						logger.Error(
							err.Error(),
							"duration", time.Since(start),
						)

						fmt.Println(debug.Stack())
					} else {
						logger.Error(
							"Unknown error",
							"duration", time.Since(start),
						)
					}
				}
			}()

			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Info(
				"request",
				"status", wrapped.status,
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}
