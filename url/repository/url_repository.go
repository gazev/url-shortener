package repository

import (
	"bytes"
	"context"
	"encoding/gob"
	"mus/url/domain"

	"github.com/go-redis/redis/v8"
)

type URLRepository struct {
	db *redis.Client
}

func NewURLRepository(db *redis.Client) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) GetUrl(hash string) (domain.URL, error) {
	b, err := r.db.Get(context.Background(), hash).Bytes()
	if err != nil {
		return domain.URL{}, err
	}

	var res domain.URL
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(&res); err != nil {
		return domain.URL{}, err
	}

	return res, nil
}

func (r *URLRepository) SetUrl(url domain.URL) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(url); err != nil {
		return err
	}

	if err := r.db.Set(context.Background(), url.Hash, b.Bytes(), 0).Err(); err != nil {
		return err
	}

	return nil
}
