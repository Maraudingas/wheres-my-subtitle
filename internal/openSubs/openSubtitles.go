package openSubs

import (
	"log/slog"

	opensubtitles "github.com/angelospk/opensubtitles-go"
)

type OpenSubs struct {
	log    *slog.Logger
	Client *opensubtitles.Client
}

func NewOpenSubsClient(l *slog.Logger, apiKey string) (*OpenSubs, error) {
	openSubsClient, err := opensubtitles.NewClient(opensubtitles.Config{ApiKey: apiKey})
	if err != nil {
		l.Error("Failed to create client to OpenSubs", "Error", err)
		return nil, err
	}
	return &OpenSubs{
		log:    l,
		Client: openSubsClient,
	}, nil
}
