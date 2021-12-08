package lanternfish

import (
	"testing"
)

func TestPopulation(t *testing.T) {
	expected := 26984457539
	got := population(256, 6, 2, "input_test.txt")
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
