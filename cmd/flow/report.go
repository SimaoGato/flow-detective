package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/SimaoGato/flow-detective/internal/data"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a text report for the current sprint",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Load Data
		path := filepath.Join(".flow", "context.yaml")
		ctx, err := data.LoadContext(path)
		if err != nil {
			fmt.Printf("❌ Error loading context: %v\n", err)
			return
		}

		// 2. Header
		fmt.Println("========================================")
		fmt.Printf("🕵️  FLOW REPORT: %s\n", ctx.ProjectName)
		fmt.Printf("📅  Date: %s\n", time.Now().Format("2006-01-02"))
		fmt.Printf("🏃  Iteration: %s\n", ctx.CurrentIteration)
		fmt.Println("========================================")
		fmt.Println()

		// 3. Iterate Stories
		totalTimeMins := 0

		if len(ctx.Stories) == 0 {
			fmt.Println("(No stories recorded. Check your .flow/context.yaml)")
		}

		for _, story := range ctx.Stories {
			fmt.Printf("## [%s] %s\n", story.ID, story.Name)

			for _, task := range story.Tasks {
				// Sum up time for this task
				taskMins := 0
				for _, entry := range task.Entries {
					taskMins += entry.DurationMins
				}
				totalTimeMins += taskMins

				// Format Duration (e.g., 90m -> 1h 30m)
				durationStr := fmt.Sprintf("%dm", taskMins)
				if taskMins >= 60 {
					h := taskMins / 60
					m := taskMins % 60
					durationStr = fmt.Sprintf("%dh %dm", h, m)
				}

				status := "🚧"
				if task.Completed {
					status = "✅"
				}

				fmt.Printf("- %s %s: **%s**\n", status, task.Name, durationStr)

				// List individual notes (logs)
				for _, entry := range task.Entries {
					if entry.DurationMins > 0 {
						fmt.Printf("  * %s (%dm)\n", entry.Note, entry.DurationMins)
					}
				}
			}
			fmt.Println()
		}

		// 4. Footer
		fmt.Println("----------------------------------------")
		fmt.Printf("∑  TOTAL LOGGED: %dh %dm\n", totalTimeMins/60, totalTimeMins%60)
		fmt.Println("========================================")
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
