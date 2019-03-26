package main

import (
	"fmt"
	"github.com/kyani-inc/TranslationEngine/controller"
	"github.com/kyani-inc/TranslationEngine/database"
	"github.com/kyani-inc/TranslationEngine/models/translation_key"
	"github.com/kyani-inc/TranslationEngine/translate"
	"net/http"
)

func main() {
	database.Init()
	controller.Routes()
	translate.InitializeAwsTranslate()

	translation_key.TK = new(translation_key.TranslationKeysList)
	translation_key.TK.List = []translation_key.TranslationKeys{}
	translation_key.TK.UpdateFromDB()

	fmt.Println("Listening on port 5000...")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
