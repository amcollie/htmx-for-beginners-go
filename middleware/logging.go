package middleware

import (
	"log"
	"net/http"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		log.Printf("status: %d, method: %s, path: %s", wrapped.statusCode, r.Method, r.URL.Path)
	}
}
