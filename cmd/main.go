package main

import (
	"log/slog"

	"github.com/Maraudingas/wheres-my-subtitle/internal"
)

func main() {
	slog.Info("Starting Wheres-My-Subtitle Service")
	name, err := internal.GetSubtitle()
	if err != nil {
		slog.Error("Failed to retrieve subtitle name.", "Error", err, "Name", name)
	}
	slog.Info("Got the name of a Video for Subtitle.", "Video", name)
}
