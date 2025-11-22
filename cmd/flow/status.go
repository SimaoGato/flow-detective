package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SimaoGato/flow-detective/internal/data"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current task health",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Load Data
		path := filepath.Join(".flow", "context.yaml")
		ctx, err := data.LoadContext(path)
		if err != nil {
			fmt.Printf("Error loading context: %v\n", err)
			return
		}

		if ctx.ActiveTaskName == "" {
			fmt.Println("No active task. Run 'flow start <name>' first.")
			return
		}

		// 2. Calculate Real Stats
		// We need to find the active task and sum its entries
		var totalMins int
		var estimateMins int

		// Find the task (logic similar to checkin.go)
		// For MVP, we search strictly by name
		found := false
		for _, story := range ctx.Stories {
			for _, task := range story.Tasks {
				if task.Name == ctx.ActiveTaskName {
					estimateMins = task.EstimateMins
					for _, entry := range task.Entries {
						totalMins += entry.DurationMins
					}
					found = true
				}
			}
		}

		// If task not found in stories (e.g. ad-hoc task), use defaults
		if !found {
			// Just use time since start if we can't find recorded entries
			sinceStart := int(time.Since(ctx.LastActivity).Minutes())
			totalMins = sinceStart
			estimateMins = 60 // Default 1 hour budget if unknown
		}

		pct := float64(totalMins) / float64(estimateMins)
		if pct > 1.0 {
			pct = 1.0
		}

		// 3. Start TUI
		p := tea.NewProgram(model{
			task:    ctx.ActiveTaskName,
			percent: pct,
			used:    totalMins,
			total:   estimateMins,
			prog:    progress.New(progress.WithDefaultGradient()),
		})

		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	},
}

// --- Bubble Tea Model ---

type model struct {
	task    string
	percent float64
	used    int
	total   int
	prog    progress.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit // Quit on any key
	case tea.WindowSizeMsg:
		m.prog.Width = msg.Width - 4
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	pad := lipgloss.NewStyle().Padding(1).Render
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4"))

	title := titleStyle.Render(fmt.Sprintf("🕵️  Task: %s", m.task))
	stats := fmt.Sprintf("%d mins / %d mins", m.used, m.total)

	bar := m.prog.ViewAs(m.percent)

	return pad(title + "\n" + bar + "\n" + stats + "\n\n(Press any key to exit)")
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
