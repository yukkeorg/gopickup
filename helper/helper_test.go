package helper

import (
	"strings"
	"testing"
)

func TestInputData(t *testing.T) {
	input := "line1\nline2\nline3\n"
	reader := strings.NewReader(input)
	expected := []string{"line1", "line2", "line3"}

	result := InputData(reader)

	if len(result) != len(expected) {
		t.Fatalf("expected %d lines, but got %d", len(expected), len(result))
	}

	for i, line := range result {
		if line != expected[i] {
			t.Errorf("expected line %d to be %q, but got %q", i, expected[i], line)
		}
	}
}
