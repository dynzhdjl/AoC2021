package dumbo_octopus

import "testing"

func TestOctopusFlashCount(t *testing.T) {
	octopus := [][]int{{1, 1, 1, 1, 1}, {1, 9, 9, 9, 1}, {1, 9, 1, 9, 1}, {1, 9, 9, 9, 1}, {1, 1, 1, 1, 1}}
	expected := 9
	steps := 1
	got := countFlashes(octopus, steps)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
	steps = 10
	expected = 204
	got = countFlashes(getOctopusMap("input_test.txt"), steps)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}

	steps = 100
	expected = 1656
	got = countFlashes(getOctopusMap("input_test.txt"), steps)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestSyncronizationStep(t *testing.T) {
	expected := 195
	got := findSyncronizationStep(getOctopusMap("input_test.txt"))
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
