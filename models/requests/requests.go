package requests

type AddTranslationInput struct {
	Key         string `json:"key"`
	Translation string `json:"translation"`
	Language    string `json:"language"`
}

type DeleteTranslationInput struct {
	Key      string `json:"key"`
	Language string `json:"language"`
}
