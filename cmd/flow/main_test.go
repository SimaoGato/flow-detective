package main

import "testing"

func TestSanity(t *testing.T) {
	// This test always passes. It proves the test runner works.
	if 1+1 != 2 {
		t.Errorf("Math is broken in the universe.")
	}
}
