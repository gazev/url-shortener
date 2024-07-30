package url

import (
	"mus/logger"
	"mus/url/domain"
	"mus/url/repository"
	"strings"
)

func CreateShortURL(r CreateShortURLRequest, ur *repository.URLRepository) (domain.URL, error) {
	u, err := domain.NewUrl(r.URL)
	if err != nil {
		return domain.URL{}, err
	}

	dbUrl, err := GetShortURL(u.Hash, ur)
	if err == nil {
		if dbUrl.URL == u.URL {
			return dbUrl, nil
		}
		logger.LogInfo("colliding hashes!!! overwritting %s from database\n", dbUrl.URL)
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
		return domain.URL{}, err
	}

	return u, nil
}
