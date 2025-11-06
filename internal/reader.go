package internal

import (
	"bufio"
	"log/slog"
	"os"
)

func GetSubtitle() (string, error) {
	slog.Debug("Reading input from os reader for subtitle name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		slog.Error("failed to read subtitle", "Error", err)
		return "", err
	}
	slog.Debug("Reading input from os finished successfuly")
	return name, nil
}
