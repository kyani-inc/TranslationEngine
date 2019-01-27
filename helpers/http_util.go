package helpers

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"net/http"
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

func GetTranslationSettingsFromHeader(r *http.Request) string  {

	acceptLanguage := r.Header.Get("Accept-Language")

	acceptLanguage = strings.Split(acceptLanguage, ",")[0]

	locale := strings.Split(acceptLanguage, "-")
	reversedlocale := reverse(locale)

	return strings.Join(reversedlocale, "-")
}

func reverse(val []string) []string {
	for i := 0; i < len(val)/2; i++ {
		j := len(val) - i - 1
		val[i], val[j] = val[j], val[i]
	}
	return val
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
