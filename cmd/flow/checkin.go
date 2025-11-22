package main

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/SimaoGato/flow-detective/internal/data"
	"github.com/SimaoGato/flow-detective/internal/logger"
	"github.com/spf13/cobra"
)

const reviewThreshold = 4 * 60 // 4 hours in minutes

var checkinCmd = &cobra.Command{
	Use:    "checkin",
	Short:  "Internal command called by git hooks",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Setup
		path := filepath.Join(".flow", "context.yaml")
		_ = logger.Setup(path)

		// 2. Load Data
		ctx, err := data.LoadContext(path)
		if err != nil {
			slog.Error("checkin failed: load error", "err", err)
			return
		}

		// 3. Calculate Duration
		now := time.Now()
		duration := now.Sub(ctx.LastActivity)
		minutes := int(duration.Minutes())

		// 4. Sanity Check
		if minutes < 1 {
			slog.Info("checkin ignored: too short", "mins", minutes)
			return
		}

		// 5. Detective Logic
		needsReview := false
		if minutes > reviewThreshold {
			needsReview = true
			slog.Info("suspicious duration detected", "mins", minutes)
		}

		// 6. Create Entry
		newEntry := data.Entry{
			Timestamp:    now,
			DurationMins: minutes,
			Note:         "Git Commit",
			NeedsReview:  needsReview,
		}

		// 7. Append to the Active Task
		// Simple search for MVP
		taskFound := false
		for i := range ctx.Stories {
			for j := range ctx.Stories[i].Tasks {
				if ctx.Stories[i].Tasks[j].Name == ctx.ActiveTaskName {
					ctx.Stories[i].Tasks[j].Entries = append(ctx.Stories[i].Tasks[j].Entries, newEntry)
					taskFound = true
					break
				}
			}
			if taskFound {
				break
			}
		}

		if !taskFound {
			slog.Warn("checkin warning: active task not found", "task", ctx.ActiveTaskName)
		}

		// 8. Save
		ctx.LastActivity = now
		if err := data.SaveContext(path, ctx); err != nil {
			slog.Error("checkin failed: save error", "err", err)
			return
		}

		// Success
		slog.Info("checkin success", "task", ctx.ActiveTaskName, "mins", minutes)
		fmt.Printf("🕵️  Logged %dm to '%s'\n", minutes, ctx.ActiveTaskName)
	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)
}
