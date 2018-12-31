package main

import (
	"github.com/catmullet/TranslationEngine/models"
	"encoding/json"
	"fmt"
)

func main() {

	newKeyMap := models.TranslationKeys{}

	newKeyMap.Language = "en"

	newKeyMap.AddKey("hello.world.whatup", "This is the key value")
	newKeyMap.AddKey("hello.world.senior", "This is the key value for senior")
	newKeyMap.AddKey("hello.world.clubbin", "This is the key value for clubbin")
	newKeyMap.AddKey("hello.world.chicca", "This is the key value chicca")
	newKeyMap.AddKey("hello.universe.hello", "Is this working?")
	newKeyMap.AddKey("default.universe.hello", "Is this working?")

	newKeyMap.DeleteKey("hello.world.chicca")

	json, _ := json.Marshal(newKeyMap)

	fmt.Println(string(json))
	fmt.Println(newKeyMap.GetKey("hello.universe.hello"))
}