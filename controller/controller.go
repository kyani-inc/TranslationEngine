package controller

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/TranslationEngine/helpers"
	"github.com/catmullet/TranslationEngine/models/translation_key"
	"net/http"
)

func AddTranslation(w http.ResponseWriter, r *http.Request) {

}

func GetTranslationsByLanguage(w http.ResponseWriter, r *http.Request) {
	lang, err := helpers.GetLanguageFromPath(r.URL)

	fmt.Println(lang)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tks := translation_key.Get(lang)
	json, _ := json.Marshal(tks.KeyMap)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(json))
}
