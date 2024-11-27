package httpserver

import (
	"fmt"
	"net/http"

	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/ctxutil"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging"
	"github.com/raphoester/clickplanet.lol-backend/internal/pkg/logging/lf"
)

func MiddlewareStack(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			handler = middlewares[i](handler)
		}
		return handler
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IPReaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.Header.Get("X-Real-IP")
		}
		if ip == "" {
			ip = r.RemoteAddr
		}
		ctx := ctxutil.AddIPToContext(r.Context(), ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewLoggingMiddleware(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapped := &wrappedWriter{ResponseWriter: w}
			next.ServeHTTP(wrapped, r)
			logger.Info(
				"new request on web server",
				lf.String("method", r.Method),
				lf.String("uri", r.RequestURI),
				lf.Int("status_code", wrapped.code),
			)
		})
	}
}

type wrappedWriter struct {
	code int
	http.ResponseWriter
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
