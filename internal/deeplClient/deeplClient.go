package deeplclient

import (
	"bytes"
	"encoding/json"
	"errors"
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

func NewDeeplClient(l *slog.Logger, apiKey string) (*DeeplClient, error) {
	if apiKey == "" {
		return nil, errors.New("Deepl Api Key is required!")
	}
	return &DeeplClient{
		log:    l,
		Client: &http.Client{},
		apiUrl: GetApiUrl(apiKey),
		apiKey: apiKey,
	}, nil
}

func GetApiUrl(apiKey string) string {
	if strings.HasSuffix(apiKey, "fx") {
		return apiUrlFree
	}
	return apiUrl
}

func (c *DeeplClient) GetTranslation(text, language string) (string, error) {
	jsonBody, err := json.Marshal(deeplclient.TranslationRequest{Text: []string{text}, TargetLang: language})
	if err != nil {
		return "", errors.New("Failed to Marshal Json Body into bytes")
	}
	req, err := http.NewRequest("POST", c.apiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", errors.New("Failed to create new HTTP request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+c.apiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", errors.New("Failed to do a HTTP request")
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to read response body from bytes")
	}
	var result deeplclient.TranslationsResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", errors.New("Failed to unmarshal json body bytes to struct")
	}
	c.log.Info("Received Translation", "Text", result.Translations[0].Text, "Language", result.Translations[0].DetectedSourceLanguage)
	return result.Translations[0].Text, nil
}
