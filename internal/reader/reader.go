package reader

import (
	"bufio"
	"fmt"
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

func (r *Reader) Read(i string, e string) string {
	r.log.Debug("Reading input from os reader")
	fmt.Print(i)
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		r.log.Error(e, "Error", err)
		os.Exit(1)
		return ""
	}
	r.log.Debug("Reading input from os finished successfuly")
	return strings.TrimSpace(name)
}
