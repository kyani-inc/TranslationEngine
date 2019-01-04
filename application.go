package main

import (
	"fmt"
	"github.com/catmullet/TranslationEngine/controller"
	"github.com/catmullet/TranslationEngine/database"
	"github.com/catmullet/TranslationEngine/models/translation_key"
	"net/http"
)

func main() {
	database.Init()
	controller.Routes()

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
