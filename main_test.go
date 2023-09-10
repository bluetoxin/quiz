package main

import (
	"strings"
	"testing"
)

func TestShuffleQuestions(t *testing.T) {
	// Create a test dataset
	lines := [][]string{
		{"Question1", "Answer1"},
		{"Question2", "Answer2"},
		{"Question3", "Answer3"},
		{"Question4", "Answer4"},
	}

	// Shuffle the questions
	shuffleQuestions(lines)

	// Ensure that the number of lines is still the same
	if len(lines) != 4 {
		t.Errorf("Expected 4 lines, got %d", len(lines))
	}

	// Check if any two consecutive lines are the same (indicating a shuffle)
	for i := 0; i < len(lines)-1; i++ {
		if strings.EqualFold(lines[i][0], lines[i+1][0]) {
			t.Errorf("Lines not shuffled")
			break
		}
	}
}
