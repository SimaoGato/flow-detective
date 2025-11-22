package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitCommand(t *testing.T) {
	// 1. Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "flow-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	// Cleanup: Delete temp dir when test finishes
	defer os.RemoveAll(tempDir)

	// 2. Save current working directory so we can go back
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current wd: %v", err)
	}

	// 3. Change directory to the temp dir
	// This ensures 'flow init' creates the .flow folder inside tempDir
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change dir: %v", err)
	}
	// Ensure we go back to the original directory when test is done
	defer func() {
		_ = os.Chdir(originalWd)
	}()

	// 4. Run the command manually
	// initCmd is defined in init.go, which is part of package main
	initCmd.Run(initCmd, []string{})

	// 5. Verify the file was created
	expectedFile := filepath.Join(".flow", "context.yaml")
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Errorf("Expected file %s was not created", expectedFile)
	}
}
