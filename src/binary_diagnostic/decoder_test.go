package binary_diagnostic

import (
	"testing"
)

func TestDecoder(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	gamma, epsilon := decode(input)
	got := gamma * epsilon
	expected := uint64(198)

	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestLifeRateDecoder(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	got := decodeLifeSupportRatings(input, o2) * decodeLifeSupportRatings(input, co2)
	expected := uint64(230)

	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
