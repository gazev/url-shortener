package api

import (
	"mus/logger"
	"mus/middleware"
	"mus/url"
	"mus/url/repository"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type MusAPI struct {
	Addr string
	DB   *redis.Client
}

func NewMusAPI(addr string, db *redis.Client) *MusAPI {
	return &MusAPI{
		Addr: addr,
		DB:   db,
	}
}

func (m *MusAPI) Run() error {
	mux := http.NewServeMux()
	ur := repository.NewURLRepository(m.DB)

	mux.Handle("GET /",
		middleware.Logging(
			url.IndexRoute()))

	mux.Handle("GET /{short}",
		middleware.Logging(
			url.GetShortURLRoute(ur)))

	mux.Handle("POST /",
		middleware.Logging(
			middleware.EnsureJson(
				url.CreateShortURLRoute(ur))))

	s := http.Server{
		Addr:    m.Addr,
		Handler: mux,
	}
	logger.Log("Listening on http://127.0.0.1:8000")
	return s.ListenAndServe()
}
