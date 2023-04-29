package service

import (
	"fmt"
)

var M = map[string]string{}

func SaveURL(url string) string {
	miniURL := randomString()

	M[miniURL] = url

	return miniURL
}

func GetURL(miniURL string) (string, error) {
	m := M[miniURL]
	if m == "" {
		return "", fmt.Errorf("don't have miniURL")
	}

	return m, nil
}
