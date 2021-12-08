package hydrothermal_venture

import (
	"testing"
)

func TestVentLineCreation(t *testing.T) {
	input := []string{
		"0,0 -> 0,8",
		"0,0 -> 25,0",
		"10,0 -> 8,0",
		"0,10 -> 0,8",
	}

	ventLines, m, n := parseRawData(input)
	expectedDir := [4][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	if m != 26 {
		t.Error()
	}
	if n != 11 {
		t.Error()
	}
	for k, line := range ventLines {

		if line.direction != expectedDir[k] {
			t.Errorf("expected: %v, got: %v", expectedDir[k], line.direction)
		}
	}
}
