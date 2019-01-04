package helpers

import (
	"errors"
	"strings"
)

func ConvertLocaleToCountryAndLanguage(locale string) (string, string, error) {

	countryAndLanguage := strings.Split(locale, "-")
	if len(countryAndLanguage) == 2 {
		return countryAndLanguage[0], countryAndLanguage[1], nil
	}

	return "", "", errors.New("Failed To Get Language")
}
