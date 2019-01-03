package helpers

import (
	"github.com/catmullet/TranslationEngine/models"
	"testing"
)

func TestCompare(t *testing.T) {
	tk1 := models.TranslationKeys{Language: "en"}
	tk2 := models.TranslationKeys{Language: "en"}

	same, err := Compare(tk1, tk2)

	if err != nil {
		t.Errorf("Failed With Error: %s", err)
	}

	if same != true {
		t.Error("Determined not the same when they are")
	}
}
