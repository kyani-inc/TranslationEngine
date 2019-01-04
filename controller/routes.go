package controller

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/", GetTranslationsByLanguage)
	http.HandleFunc("/add", AddTranslationByLanguage)
	http.HandleFunc("/delete", DeleteTranslationByLanguage)
}
