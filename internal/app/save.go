package app

var M = map[string]string{}

func SaveUrls(url, miniUrl string) {
	miniUrl = "/" + miniUrl
	M[miniUrl] = url
}

func GetUrl(miniUlr string) string {
	if miniUlr != "" {
		_, ok := M[miniUlr]
		if ok {
			return M[miniUlr]
		}
	}

	return ""
}
