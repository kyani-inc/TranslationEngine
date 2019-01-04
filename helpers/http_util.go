package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func GetLanguageFromPath(url *url.URL) (string, error) {
	pathParams := strings.Split(url.Path, "/")

	fmt.Println(pathParams)

	if len(pathParams) > 1 {
		lang := strings.Replace(pathParams[1], " ", "", -1)
		if len(lang) == 2 {
			return lang, nil
		}
	}

	return "", errors.New("Failed To Get Language")
}
