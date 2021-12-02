package sonar_sweep

import (
	"testing"
)

func TestScanner(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	got := scanner(input)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestSlidingScanner(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5
	got := slidingScanner(input)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
