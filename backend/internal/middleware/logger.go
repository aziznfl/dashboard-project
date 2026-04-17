package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// Logger is a middleware that logs the start and end of an HTTP request,
// as well as the latency and status code.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		defer func() {
			log.Printf("%s | %s | %d | %s | %s", r.Method, r.URL.Path, ww.Status(), time.Since(start).String(), r.RemoteAddr)
		}()

		next.ServeHTTP(ww, r)
	})
}
