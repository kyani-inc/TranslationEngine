package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func GetLanguageAndCountryFromPath(url *url.URL) (string, string, error) {
	pathParams := strings.Split(url.Path, "/")

	fmt.Println(pathParams)

	if len(pathParams) > 1 {
		locale := strings.Replace(pathParams[1], " ", "", -1)
		if len(locale) == 5 {
			countryAndLanguage := strings.Split(locale, "-")
			if len(countryAndLanguage) == 2 {
				return countryAndLanguage[0], countryAndLanguage[1], nil
			}
		}
	}

	return "", "", errors.New("Failed To Get Language")
}

func GetLocaleFromPath(url *url.URL) (string, bool, error) {
	isJS := false

	pathParams := strings.Split(url.Path, "/")

	fmt.Println(pathParams)

	if len(pathParams) > 1 {
		locale := strings.ToLower(pathParams[1])
		locale = strings.Replace(locale, " ", "", -1)
		if len(locale) == 5 {
			return locale, isJS, nil
		}
		if len(locale) == 8 {
			locale = strings.Replace(locale, ".js", "", -1)
			isJS = true
			return locale, isJS, nil
		}
	}

	return "", isJS, errors.New("Failed To Get Language")
}
