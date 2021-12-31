package transparent_origami

import (
	"fmt"
	"testing"
)

func TestFold(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	s, e := [2]int{0, 0}, [2]int{len(grid) - 1, len(grid[0]) - 1}
	s, e = foldAlongY(s, e, grid, 7)
	for k := range grid {
		fmt.Println(grid[k])
	}
	expected := 17
	got := count(s, e, grid)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
