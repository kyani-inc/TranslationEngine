package requests

type AddTranslationInput struct {
	Key         string `json:"key"`
	Translation string `json:"translation"`
	Locale      string `json:"locale"`
}

type DeleteTranslationInput struct {
	Key    string `json:"key"`
	Locale string `json:"locale"`
}

type TranslateInput struct {
	Key          string `json:"key"`
	SourceLocale string `json:"source_locale"`
	TargetLocale string `json:"target_locale"`
	Text         string `json:"text"`
	CreateKey    bool   `json:"create_key"`
}

type SyncInput struct {
	SourceLocale string `json:"source_locale"`
	TargetLocale string `json:"target_locale"`
}
