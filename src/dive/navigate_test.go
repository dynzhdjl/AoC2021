package dive

import (
	"testing"
)

func TestNavigation(t *testing.T) {
	input := []command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2}}

	expected := 150
	got := navigate(input, updateWithOutAim)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestNavigationWithAim(t *testing.T) {
	input := []command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2}}

	expected := 900
	got := navigate(input, updateWithAim)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
