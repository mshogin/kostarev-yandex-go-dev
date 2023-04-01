package app

var M = map[string]string{}

func SaveUrls(url, miniURL string) {
	miniURL = "/" + miniURL
	M[miniURL] = url
}

func GetURL(miniURL string) string {
	if miniURL != "" {
		_, ok := M[miniURL]
		if ok {
			return M[miniURL]
		}
	}

	return ""
}
