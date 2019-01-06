package controller

import (
	"encoding/json"
	"fmt"
	"github.com/catmullet/TranslationEngine/helpers"
	"github.com/catmullet/TranslationEngine/models/requests"
	"github.com/catmullet/TranslationEngine/models/translation_key"
	"github.com/catmullet/TranslationEngine/translate"
	"io/ioutil"
	"net/http"
)

func AddTranslationByLanguage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

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

		tk := translation_key.Get(ati.Locale)

		if tk.IsEmpty() {
			tk.Locale = ati.Locale
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

func GetTranslationsByLocale(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	locale, err := helpers.GetLocaleFromPath(r.URL)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tks := translation_key.Get(locale)
	json, _ := json.Marshal(tks.KeyMap)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(json))
}

func DeleteTranslationByLocale(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

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

	tks := translation_key.Get(dti.Locale)
	tks.DeleteKey(dti.Key)
	err = translation_key.Put(tks)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "Successfully Deleted Translation Key")
}

func TranslateKeyToLanguage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		ti := requests.TranslateInput{}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = json.Unmarshal(body, &ti)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Return
		if ti.SourceLocale != "" && ti.TargetLocale != "" && (ti.Text != "" || ti.Key != "") {
			translation := translate.ParseTranslateRequest(ti)
			fmt.Fprintf(w, "%s", translation)
			return
		}

		break
	default:
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed %s", r.Method)
		return
	}

	fmt.Fprintf(w, "Missing required parameters")
}

func ListCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	body, _ := json.Marshal(translation_key.GetAllCountriesAndLanguages())

	fmt.Fprintf(w, "%s", string(body))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
