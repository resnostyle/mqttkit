// Package logx configures process-wide slog logging from a level string.
package logx

import (
	"log/slog"
	"os"
	"strings"
)

// Configure sets the default slog handler writing to stdout.
// When json is true, logs are JSON; otherwise text.
func Configure(level string, json bool) {
	var lvl slog.Level
	switch strings.ToUpper(level) {
	case "DEBUG":
		lvl = slog.LevelDebug
	case "WARN", "WARNING":
		lvl = slog.LevelWarn
	case "ERROR":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{Level: lvl}
	if json {
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))
		return
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
}
