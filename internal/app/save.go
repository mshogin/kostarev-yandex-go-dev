package app

var M = map[string][]byte{}

func SaveUrls(url []byte, miniURL string) {
	miniURL = "/" + miniURL
	M[miniURL] = url
}

func GetURL(miniURL string) []byte {
	if miniURL != "" {
		_, ok := M[miniURL]
		if ok {
			return M[miniURL]
		}
	}

	return []byte("")
}
