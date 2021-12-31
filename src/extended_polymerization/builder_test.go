package extended_polymerization

import (
	"testing"
)

func TestRuleCreation(t *testing.T) {
	input := []string{"CH -> B", "HH -> N", "CB -> H", "HB -> B"}
	rules := makeRules(input)
	expected := 4
	got := len(rules)
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
func TestInputParsing(t *testing.T) {
	template, rules := parseRawData("input_test.txt")
	expectedTemplate := "NNCB"
	expectedRuleCount := 16

	if expectedRuleCount != len(rules) {
		t.Errorf("Rules parsing failed")
	}
	if template != expectedTemplate {
		t.Errorf("template parsing failed")
	}
}

func TestBuild(t *testing.T) {
	expected := uint64(1588)
	template, rules := parseRawData("input_test.txt")
	elements := build(template, rules, 10)
	min, max := minMax(elements)
	minCnt, _ := elements[min]
	maxCnt, _ := elements[max]
	got := maxCnt - minCnt
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
