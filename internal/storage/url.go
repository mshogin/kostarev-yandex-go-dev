package storage

import (
	"errors"
)

var M = map[string]string{}

func SaveURL(url string) string {
	miniURL := "/" + randomString() // необходимо добавлять / потому что потом без него не находится элемент

	M[miniURL] = url

	return miniURL
}

func GetURL(miniURL string) (string, error) {
	m := M[miniURL]
	if m == "" {
		return "", errors.New("don't have miniURL")
	}

	return m, nil
}
