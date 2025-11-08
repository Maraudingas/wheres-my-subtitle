package deeplclient

type TranslationRequest struct {
	Text       []string `json:"text"`
	TargetLang string   `json:"target_lang"`
}
