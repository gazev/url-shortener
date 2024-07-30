package url

import (
	"encoding/json"
	"errors"
	"mus/logger"
	"mus/url/repository"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func IndexRoute() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<i>\"To be fond of dancing was a certain step towards falling in love\"</i>"))
	})
}

func CreateShortURLRoute(ur *repository.URLRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req CreateShortURLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.LogDebug("invalid request -> %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		resp, err := CreateShortURL(req, ur)
		if err != nil {
			logger.LogDebug("failed creating url -> %s\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			logger.LogDebug("failed sending response -> %s\n", err)
			w.WriteHeader(http.StatusOK)
		}
	})
}

func GetShortURLRoute(ur *repository.URLRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := r.URL.Path[1:]
		url, err := GetShortURL(hash, ur)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			logger.LogDebug("failed getting url -> %s\n", err)
			http.Error(w, "Failed processing request", http.StatusInternalServerError)
			return
		}

		if hash[len(hash)-1] == '+' {
			if err := json.NewEncoder(w).Encode(url); err != nil {
				logger.LogDebug("failed encoding get url response -> %s\n", url)
				http.Error(w, "Failed processing request", http.StatusInternalServerError)
			}
			return
		}

		http.Redirect(w, r, url.URL, http.StatusFound)
	})
}
