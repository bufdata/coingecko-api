package util

import (
	"log/slog"
	"os"
	"path/filepath"
)

// GetLogger returns default customized log.
func GetLogger(name string) {
	replacer := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replacer}).WithAttrs(
		[]slog.Attr{slog.String("APIName", name)})
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)
}
