package reader

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

type Reader struct {
	log *slog.Logger
}

func NewReader(l *slog.Logger) *Reader {
	return &Reader{
		log: l,
	}
}

func (r *Reader) Read() (string, error) {
	r.log.Debug("Reading input from os reader")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		r.log.Error("failed to read input", "Error", err)
		return "", err
	}
	r.log.Debug("Reading input from os finished successfuly")
	return strings.TrimSpace(name), nil
}
