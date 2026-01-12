package storage

import "strings"

var (
	urlStore = make(map[string]string)
	counter  = 1000000
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func encode(n int) string {
	if n == 0 {
		return string(charset[0])

	}
	var sb strings.Builder
	for n > 0 {
		sb.WriteByte(charset[n%len(charset)])
		n /= len(charset)
	}
	return sb.String()
}

// function for saving the original url corresponding to the shorten url and return  the shorten url
func SaveURL(originalURL string) string {
	shortID := encode(counter)
	urlStore[shortID] = originalURL
	counter++
	return shortID
}

// function for getting the original
func GetURL(shortenURL string) (string, bool) {
	url, exist := urlStore[shortenURL]
	return url, exist
}
