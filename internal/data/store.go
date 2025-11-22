package data

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadContext reads the yaml file from the given path
func LoadContext(path string) (*Context, error) {
	// 1. Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("context file not found at %s", path)
	}

	// 2. Read bytes
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 3. Unmarshal (Parse) YAML into Struct
	var ctx Context
	if err := yaml.Unmarshal(data, &ctx); err != nil {
		return nil, fmt.Errorf("invalid yaml format: %w", err)
	}

	return &ctx, nil
}

// SaveContext writes the struct back to the file
func SaveContext(path string, ctx *Context) error {
	// 1. Create directory if it doesn't exist (e.g., .flow/)
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 2. Marshal (Convert) Struct to YAML bytes
	data, err := yaml.Marshal(ctx)
	if err != nil {
		return err
	}

	// 3. Write to file (0644 = readable by user)
	return os.WriteFile(path, data, 0644)
}
