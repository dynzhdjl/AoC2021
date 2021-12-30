package transparent_origami

import (
	"fmt"
	"testing"
)

func TestFold(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0}, //0
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, //1
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //2
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //3
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1}, //4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //5
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //6
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //7
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //8
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //9
		{0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0}, //10
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, //11
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1}, //12
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //13
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, //14
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
	fmt.Println("=====")
	s, e = foldAlongX(s, e, grid, 5)
	for k := range grid {
		fmt.Println(grid[k])
	}

}
