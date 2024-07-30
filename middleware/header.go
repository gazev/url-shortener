package middleware

import "net/http"

func EnsureJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := r.Header.Get("Content-Type")
		if content != "application/json" {
			http.Error(w, "Bad Request", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
