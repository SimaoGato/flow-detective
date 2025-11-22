package data

import (
	"os"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// 1. Define a dummy context
	original := &Context{
		ProjectName:      "Flow Detective",
		CurrentIteration: "Sprint 1",
		Stories: []Story{
			{
				ID:   "REV-101",
				Name: "Setup",
			},
		},
	}

	// 2. Define a temp file path
	tmpFile := "test_context.yaml"
	defer os.Remove(tmpFile) // Cleanup after test

	// 3. Test Save
	err := SaveContext(tmpFile, original)
	if err != nil {
		t.Fatalf("Failed to save: %v", err)
	}

	// 4. Test Load
	loaded, err := LoadContext(tmpFile)
	if err != nil {
		t.Fatalf("Failed to load: %v", err)
	}

	// 5. Verify Data
	if loaded.ProjectName != original.ProjectName {
		t.Errorf("Expected %s, got %s", original.ProjectName, loaded.ProjectName)
	}
}
