package middleware

import "net/http"

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		content := r.Header.Get("Authorization")
		if content != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}
