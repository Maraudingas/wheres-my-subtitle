package deeplclient

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Maraudingas/wheres-my-subtitle/api/types/deeplclient"
)

const apiURL = "https://api-free.deepl.com/v2/translate"

func TestClient() {
	apiKey := os.Getenv("API_KEY")
	jsonBody, err := json.Marshal(deeplclient.TranslationRequest{Text: []string{"Hello World!"}, TargetLang: "LT"})
	if err != nil {
		log.Panic("failed to marshal")
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Panic("Failed to Create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic("Failed to post a request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status not eq HTTP OK")
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Failed to read response Body to Bytes")
	}
	var result deeplclient.TranslationsResponse
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		log.Fatal("Failed to unmarshal response from bytes")
	}
	log.Printf("Text:%v, Language:%s", result.Translations[0].Text, result.Translations[0].DetectedSourceLanguage)
}
