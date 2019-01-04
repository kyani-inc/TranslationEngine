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

func GetLocaleFromPath(url *url.URL) (string, error) {
	pathParams := strings.Split(url.Path, "/")

	fmt.Println(pathParams)

	if len(pathParams) > 1 {
		locale := strings.Replace(pathParams[1], " ", "", -1)
		if len(locale) == 5 {
			return locale, nil
		}
	}

	return "", errors.New("Failed To Get Language")
}
