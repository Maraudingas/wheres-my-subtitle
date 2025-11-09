package main

import (
	"context"
	"fmt"
	"strings"

	deeplclient "github.com/Maraudingas/wheres-my-subtitle/internal/deeplClient"
	"github.com/Maraudingas/wheres-my-subtitle/internal/logging"
	"github.com/Maraudingas/wheres-my-subtitle/internal/openSubs"
	"github.com/Maraudingas/wheres-my-subtitle/internal/reader"
	"github.com/angelospk/opensubtitles-go"
)

func main() {

	logger := logging.NewLogger()

	reader := reader.NewReader(logger)

	deeplApiKey := reader.Read("Please write deepl API Key: ", "Failed to retrieve deepl API Key.")

	deeplClient, err := deeplclient.NewDeeplClient(logger, deeplApiKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Deepl Client, Error: %v", err))
	}

	translation, err := deeplClient.GetTranslation("This is the Best Program Ever written!", "LT")
	if err != nil {
		panic(fmt.Sprintf("Failed to translate, Error: %v", err))
	}
	logger.Info("Let's see what we got translated", "Translation", translation)

	subtitle := reader.Read("Please Write Video name for subtitle search: ", "Failed to retrieve subtitle name.")

	logger.Info("Got the name of a Video for Subtitle.", "Video", strings.TrimSpace(subtitle))

	openSubApiKey := reader.Read("Please Write API key for OpenSubs access: ", "Failed to retrieve API name.")

	opensubby, err := openSubs.NewOpenSubsClient(logger, openSubApiKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize OpenSubClient, Error: %v", err))
	}
	logger.Info("current opensubs website", "URL", opensubby.Client.GetCurrentBaseURL())

	respone, err := opensubby.Client.SearchSubtitles(context.TODO(), opensubtitles.SearchSubtitlesParams{Query: &subtitle})
	if err != nil {
		panic(fmt.Sprintf("Failed to Search OpenSubClient, Error: %v", err))
	}
	logger.Info("page of subtitle", "URL", respone.Data[0].Attributes.FeatureDetails.MovieName)
}
