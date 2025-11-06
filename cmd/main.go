package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Maraudingas/wheres-my-subtitle/internal/logging"
	"github.com/Maraudingas/wheres-my-subtitle/internal/openSubs"
	"github.com/Maraudingas/wheres-my-subtitle/internal/reader"
	"github.com/angelospk/opensubtitles-go"
)

func main() {
	logger := logging.NewLogger()

	reader := reader.NewReader(logger)

	fmt.Print("Please Write Video name for subtitle search:")
	subtitle, err := reader.Read()
	if err != nil {
		logger.Error("Failed to retrieve subtitle name.", "Error", err, "Name", subtitle)
		os.Exit(1)
	}
	logger.Info("Got the name of a Video for Subtitle.", "Video", strings.TrimSpace(subtitle))

	fmt.Print("Please Write API key for OpenSubs access:")
	apiKey, err := reader.Read()
	if err != nil {
		logger.Error("Failed to retrieve API name.", "Error", err, "Name", subtitle)
		os.Exit(1)
	}
	opensubby, err := openSubs.NewOpenSubsClient(logger, apiKey)
	if err != nil {
		logger.Error("Failed to initialize OpenSubClient", "Error", err)
		os.Exit(1)
	}
	logger.Info("current opensubs website", "URL", opensubby.Client.GetCurrentBaseURL())

	respone, err := opensubby.Client.SearchSubtitles(context.TODO(), opensubtitles.SearchSubtitlesParams{Query: &subtitle})
	if err != nil {
		logger.Error("Failed to Search OpenSubClient", "Error", err)
		os.Exit(1)
	}
	logger.Info("page of subtitle", "URL", respone.Data[0].Attributes.FeatureDetails.MovieName)
}
