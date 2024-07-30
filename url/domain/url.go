package domain

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	urllib "net/url"
	"time"

	"github.com/jxskiss/base62"
)

type URL struct {
	URL     string    `json:"url"`
	Hash    string    `json:"hash"`
	Created time.Time `json:"created at"`
}

func NewUrl(url string) (URL, error) {
	if url == "" {
		return URL{}, fmt.Errorf("empty URL")
	}

	if !isValidUrl(url) {
		return URL{}, fmt.Errorf("invalid URL format %s", url)
	}

	return URL{
		URL:     url,
		Hash:    hashUrl(url),
		Created: time.Now(),
	}, nil
}

func isValidUrl(url string) bool {
	u, err := urllib.ParseRequestURI(url)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

func hashUrl(url string) string {
	hash := crc32.ChecksumIEEE([]byte(url))
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, hash)
	return base62.EncodeToString(buf)
}
