package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"golang.org/x/exp/slog"
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
			defer func() {
				if recovered := recover(); recovered != nil {

					err, ok := recovered.(error)

					w.WriteHeader(http.StatusInternalServerError)
					if ok {
						logger.Error(err.Error())
						fmt.Println(debug.Stack())
					} else {
						logger.Error("Unknown error")
					}
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Info(
				"request",
				"status", wrapped.status,
				"method", r.Method,
				"path", r.URL.EscapedPath(),
				"duration", time.Since(start),
			)
		}

		return http.HandlerFunc(fn)
	}
}
