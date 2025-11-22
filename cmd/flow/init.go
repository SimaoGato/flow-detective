package main

import (
	"fmt"
	"log/slog" // Import the new standard logger
	"path/filepath"

	"github.com/SimaoGato/flow-detective/internal/data"
	"github.com/SimaoGato/flow-detective/internal/logger" // Import your new package
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new flow context",
	Run: func(cmd *cobra.Command, args []string) {
		path := filepath.Join(".flow", "context.yaml")

		// 1. Setup the Logger (New Step)
		// If logger fails, we just print a warning, don't crash
		if err := logger.Setup(path); err != nil {
			fmt.Printf("⚠️  Warning: Could not setup logger: %v\n", err)
		}

		defaultCtx := &data.Context{
			ProjectName:      "New Project",
			CurrentIteration: "Sprint 1",
			Stories:          []data.Story{},
		}

		fmt.Printf("🕵️  Initializing Flow Detective at %s...\n", path)

		if err := data.SaveContext(path, defaultCtx); err != nil {
			// Log the error structure for debugging
			slog.Error("failed to save context", "error", err, "path", path)
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

		// Log success silently
		slog.Info("flow initialized", "project", defaultCtx.ProjectName, "path", path)

		fmt.Println("✅ Success! You are ready to start tracking.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
