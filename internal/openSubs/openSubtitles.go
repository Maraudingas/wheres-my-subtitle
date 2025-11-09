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
	l.Debug("Creating Client to OpenSubs.")
	openSubsClient, err := opensubtitles.NewClient(opensubtitles.Config{ApiKey: apiKey})
	if err != nil {
		return nil, err
	}
	l.Debug("Successfully Created OpenSubs Client.")
	return &OpenSubs{
		log:    l,
		Client: openSubsClient,
	}, nil
}
