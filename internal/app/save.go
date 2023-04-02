package app

var M = map[string]string{}

func SaveUrls(url string, miniURL string) {
	miniURL = "/" + miniURL
	M[miniURL] = url
}

func GetURL(miniURL string) string {
	m := M[miniURL]
	return m
}
