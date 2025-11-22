package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

// Setup configures the global logger to write to a file instead of the terminal.
func Setup(storagePath string) error {
	// 1. Define log file path: .flow/debug.log
	dir := filepath.Dir(storagePath)
	logFile := filepath.Join(dir, "debug.log")

	// --- FIX START ---
	// Ensure the directory exists before creating the file
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	// --- FIX END ---

	// 2. Open the file (Append mode, Create if not exists)
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 3. Create a JSON Handler
	handler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// 4. Set as Global Logger
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return nil
}
