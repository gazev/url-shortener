package url

import (
	"errors"
	"mus/url/domain"
	"mus/url/repository"
	"strings"

	"github.com/go-redis/redis/v8"
)

func CreateShortURL(r CreateShortURLRequest, ur *repository.URLRepository) (domain.URL, error) {
	u, err := domain.NewUrl(r.URL)
	if err != nil {
		return domain.URL{}, err
	}

	dbUrl, err := GetShortURL(u.Hash, ur)
	if err == nil {
		return dbUrl, nil
	}

	if err := ur.SetUrl(u); err != nil {
		return domain.URL{}, err
	}

	return u, nil
}

func GetShortURL(hash string, ur *repository.URLRepository) (domain.URL, error) {
	hashT := strings.TrimSuffix(hash, "+")

	u, err := ur.GetUrl(hashT)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.URL{}, err
		}
		return domain.URL{}, err
	}

	return u, nil
}
