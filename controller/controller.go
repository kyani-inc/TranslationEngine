package controller

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/TranslationEngine/helpers"
	"github.com/catmullet/TranslationEngine/models/requests"
	"github.com/catmullet/TranslationEngine/models/translation_key"
	"io/ioutil"
	"net/http"
)

func AddTranslationByLanguage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		ati := requests.AddTranslationInput{}

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.Unmarshal(body, &ati)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tk := translation_key.Get(ati.Language)

		if tk.IsEmpty() {
			tk.Language = ati.Language
			tk.AddKey(ati.Key, ati.Translation)
			err = translation_key.Put(tk)

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
		}

		tk.AddKey(ati.Key, ati.Translation)
		err = translation_key.Put(tk)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		break
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed %s", r.Method)
		return
	}

	fmt.Fprintf(w, "Successfully Added Translation Key")
}

func GetTranslationsByLanguage(w http.ResponseWriter, r *http.Request) {
	lang, err := helpers.GetLanguageFromPath(r.URL)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tks := translation_key.Get(lang)
	json, _ := json.Marshal(tks.KeyMap)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(json))
}

func DeleteTranslationByLanguage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		dti := requests.DeleteTranslationInput{}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.Unmarshal(body, &dti)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tks := translation_key.Get(dti.Language)
		tks.DeleteKey(dti.Key)
		err = translation_key.Put(tks)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		break
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed %s", r.Method)
		return
	}

	fmt.Fprintf(w, "Successfully Deleted Translation Key")
}
