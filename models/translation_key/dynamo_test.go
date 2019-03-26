package translation_key

import (
	"github.com/kyani-inc/TranslationEngine/database"
	"testing"
)

func Init() {
	database.Init()
	TK = new(TranslationKeysList)
	TK.List = []TranslationKeys{}
}

func TestPut(t *testing.T) {
	Init()

	tk := TranslationKeys{}
	tk.Locale = "en"
	tk.AddKey("ui.support.error", "An error occurred")
	tk.AddKey("ui.support.info", "Info has occured")
	tk.AddKey("common.button.submit", "Submit")

	TK.AddTranslationKeys(tk).Sync()
}
