package controller

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/", GetTranslationsByLanguage)
	http.HandleFunc("/Add", AddTranslation)

}
