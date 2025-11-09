package deeplclient

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Maraudingas/wheres-my-subtitle/api/types/deeplclient"
)

const apiUrlFree = "https://api-free.deepl.com/v2/translate"
const apiUrl = "https://api.deepl.com/v2/translate"

type DeeplClient struct {
	log    *slog.Logger
	Client *http.Client
	apiUrl string
	apiKey string
}

func NewDeeplClient(l *slog.Logger, apiKey string) *DeeplClient {
	return &DeeplClient{
		log:    l,
		Client: &http.Client{},
		apiUrl: GetApiUrl(apiKey),
		apiKey: apiKey,
	}
}

func GetApiUrl(apiKey string) string {
	if strings.HasSuffix(apiKey, "fx") {
		return apiUrlFree
	} else {
		return apiUrl
	}
}

func (c *DeeplClient) GetTranslation(text, language string) string {
	jsonBody, err := json.Marshal(deeplclient.TranslationRequest{Text: []string{text}, TargetLang: language})
	if err != nil {
		c.log.Error("Failed to Marshal to Json", "Error", err)
	}
	req, err := http.NewRequest("POST", c.apiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.log.Error("Failed to create POST request", "Error", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.apiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		c.log.Error("Failed to do a HTTP request", "Error", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.log.Error("Failed to read response body", "Error", err)
	}
	var result deeplclient.TranslationsResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		c.log.Error("Failed to unmarshal json response body", "Error", err)
	}
	c.log.Info("Received Translation", "Text", result.Translations[0].Text, "Language", result.Translations[0].DetectedSourceLanguage)
	return result.Translations[0].Text
}
