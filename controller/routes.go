package controller

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/", GetTranslationsByLocale)
	http.HandleFunc("/add", AddTranslationByLanguage)
	http.HandleFunc("/delete", DeleteTranslationByLocale)

	http.HandleFunc("/translate", TranslateKeyToLanguage)
}
