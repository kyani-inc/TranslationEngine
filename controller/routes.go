package controller

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/", GetTranslations)
	http.HandleFunc("/Add", AddTranslation)

}
