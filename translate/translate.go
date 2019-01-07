package translate

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/catmullet/TranslationEngine/helpers"
	"github.com/catmullet/TranslationEngine/models/requests"
	"github.com/catmullet/TranslationEngine/models/translation_key"
)

var (
	Trans *translate.Translate
)

func InitializeAwsTranslate() {
	sess := session.Must(session.NewSession())
	sess.Config.Region = aws.String("us-east-1")
	// Create a Translate client with additional configuration
	Trans = translate.New(sess)
}

func TranslateText(text, sourceLang, targetLang string) string {

	textInput := translate.TextInput{}
	textInput.SetSourceLanguageCode(sourceLang)
	textInput.SetTargetLanguageCode(targetLang)
	textInput.SetText(text)

	to, err := Trans.Text(&textInput)

	if err != nil {
		fmt.Println(err)
	}

	return *to.TranslatedText
}

func SyncFrom(ky, from_ky translation_key.TranslationKeys) {
	for key1, _ := range from_ky.KeyMap {
		for key2, _ := range from_ky.KeyMap[key1] {
			for key3, text := range from_ky.KeyMap[key1][key2] {
				if _, ok := ky.KeyMap[key1][key2][key3]; !ok {
					_, fromLang, _ := helpers.ConvertLocaleToCountryAndLanguage(from_ky.Locale)
					_, Lang, _ := helpers.ConvertLocaleToCountryAndLanguage(ky.Locale)
					transText := TranslateText(text, fromLang, Lang)
					ky.AddKey(fmt.Sprintf("%s.%s.%s", key1, key2, key3), transText)
				}
			}
		}
	}
	translation_key.Put(ky)
}

func ParseTranslateRequest(ti requests.TranslateInput) string {

	// If they want to get the text from a key and source
	// language this will set the text based on the key and source language provided
	if ti.Key != "" && ti.SourceLocale != "" && ti.Text == "" {
		tk := translation_key.Get(ti.SourceLocale)
		ti.Text = tk.GetKey(ti.Key)
	}

	_, sourceLang, _ := helpers.ConvertLocaleToCountryAndLanguage(ti.SourceLocale)
	_, targetLang, _ := helpers.ConvertLocaleToCountryAndLanguage(ti.TargetLocale)

	translatedText := TranslateText(ti.Text, sourceLang, targetLang)

	// If CreateKey is true we will create a new key in the target language
	// copied from the source languages key
	if ti.CreateKey && ti.SourceLocale != "" && ti.TargetLocale != "" {
		tk := translation_key.Get(ti.TargetLocale)

		if tk.IsEmpty() {
			tk.Locale = ti.TargetLocale
		}

		if ti.Text != "" {
			tk.AddKey(ti.Key, ti.Text)
		} else {
			tk.AddKey(ti.Key, translatedText)
		}

		err := translation_key.Put(tk)

		// TODO: Return error back to controller here
		if err != nil {
			fmt.Println(err)
		}
	}

	return translatedText
}
