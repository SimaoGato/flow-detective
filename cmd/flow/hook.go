package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Install the git post-commit hook",
	Run: func(cmd *cobra.Command, args []string) {
		exe, _ := os.Executable()
		// The script calls our binary with 'checkin'
		scriptContent := fmt.Sprintf(`#!/bin/sh
"%s" checkin
`, exe)

		hookPath := filepath.Join(".git", "hooks", "post-commit")

		err := os.WriteFile(hookPath, []byte(scriptContent), 0755)
		if err != nil {
			fmt.Printf("❌ Error installing hook: %v\n", err)
			return
		}

		fmt.Println("🔗 Git hook installed! Flow Detective is now watching.")
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
