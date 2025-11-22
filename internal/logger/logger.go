package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

// Setup configures the global logger to write to a file instead of the terminal.
// This is crucial for the Git Hook, which runs silently.
func Setup(storagePath string) error {
	// 1. Define log file path: .flow/debug.log
	logFile := filepath.Join(filepath.Dir(storagePath), "debug.log")

	// 2. Open the file (Append mode, Create if not exists)
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 3. Create a JSON Handler (structured logging)
	// We write to the file, not os.Stdout
	handler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// 4. Set as Global Logger
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}
