package passage_pathing

import (
	"fmt"
	"testing"
)

func TestPathFinder(t *testing.T) {
	g := map[string]([]string){
		"start": []string{"A", "b"},
		"A":     []string{"start", "c", "b", "end"},
		"b":     []string{"A", "start", "d", "end"},
		"c":     []string{"A"},
		"end":   []string{"A", "b"},
		"d":     []string{"b"},
	}

	visited := map[string]bool{}
	expected := 10
	got := getNumberOfPaths(g, "start", "end", visited)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestAdjancencyList(t *testing.T) {
	input := []string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}
	g := createAdjancenyList(input)
	fmt.Println(g)
}
