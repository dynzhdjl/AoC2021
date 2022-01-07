package chiton

import "testing"

func TestRiskCalculator(t *testing.T) {
	riskMap := getRiskMap("input_test.txt")
	expected := 40
	got := calculator(riskMap)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
