package helpers

import (
	"fmt"
	"github.com/go-siris/siris/core/errors"
	"net/url"
	"strings"
)

func GetLanguageFromPath(url *url.URL) (string, error) {
	err := errors.New("")
	pathParams := strings.Split(url.Path, "/")

	fmt.Println(pathParams)

	if len(pathParams) > 1 {
		lang := strings.Replace(pathParams[1], " ", "", -1)
		if len(lang) == 2 {
			return lang, err
		}
	}

	return "", errors.New("Failed To Get Language")
}
