package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrapper struct {
	http.ResponseWriter
	status int
}

func (w *wrapper) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wr := &wrapper{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			start := time.Now()
			next.ServeHTTP(wr, r)
			log.Printf("%s: %s %d %s %d", r.RemoteAddr, r.Method, wr.status, r.URL.Path, time.Since(start))
		})
}
