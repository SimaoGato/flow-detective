package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
	"time"

	"github.com/SimaoGato/flow-detective/internal/data"
	"github.com/SimaoGato/flow-detective/internal/logger"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [task name]",
	Short: "Start working on a specific task",
	Args:  cobra.MinimumNArgs(1), // Requires at least 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Combine args into a single string
		taskName := strings.Join(args, " ")

		// 2. Setup Logger & Load Context
		path := filepath.Join(".flow", "context.yaml")
		_ = logger.Setup(path)

		ctx, err := data.LoadContext(path)
		if err != nil {
			fmt.Printf("❌ Error: Could not load context. Did you run 'flow init'?\n%v\n", err)
			return
		}

		// 3. Update Short-Term Memory
		ctx.ActiveTaskName = taskName
		ctx.LastActivity = time.Now()

		// 4. Save
		if err := data.SaveContext(path, ctx); err != nil {
			slog.Error("failed to start task", "error", err)
			fmt.Printf("❌ Error saving context: %v\n", err)
			return
		}

		// 5. Success Output
		slog.Info("task started", "task", taskName)
		fmt.Printf("⏱️  Timer started for: \033[1m%s\033[0m\n", taskName)
		fmt.Println("   (Go write some code! I'll listen for your commit.)")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
