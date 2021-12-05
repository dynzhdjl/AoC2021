package binary_diagnostic

import (
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
)

type LifeSupportDecodeType int

const (
	o2 LifeSupportDecodeType = iota
	co2
)

func decode(input []string) (uint64, uint64) {
	n := len(input)
	bitSize := len(input[0])
	gammaBinaryString := ""
	for i := 0; i < bitSize; i++ {
		onesCount := 0
		for j := 0; j < n; j++ {
			if input[j][i] == '1' {
				onesCount++
			}
		}
		if onesCount > n-onesCount {
			gammaBinaryString += "1"
		} else {
			gammaBinaryString += "0"
		}
	}
	gamma, _ := strconv.ParseUint(gammaBinaryString, 2, bitSize)
	mask := ^uint64(0) >> (64 - uint64(bitSize))
	epsilon := gamma ^ mask
	return gamma, epsilon
}

func decodeLifeSupportRatings(input []string, t LifeSupportDecodeType) uint64 {
	bitSize := len(input[0])
	values := make([]string, len(input))
	copy(values, input)
	for i := 0; i < bitSize; i++ {
		ones := []string{}
		zeros := []string{}
		for j := 0; j < len(values); j++ {
			if values[j][i] == '1' {
				ones = append(ones, values[j])
			} else if values[j][i] == '0' {
				zeros = append(zeros, values[j])
			}
		}
		switch t {
		case o2:
			if len(ones) >= len(zeros) {
				values = ones
			} else {
				values = zeros
			}
		case co2:
			if len(ones) >= len(zeros) {
				values = zeros
			} else {
				values = ones
			}
		}
		if len(values) == 1 {
			break
		}
	}
	v, _ := strconv.ParseUint(values[0], 2, bitSize)
	return v
}

func GetEnergyConsumption() uint64 {
	input := util.Read("binary_diagnostic/input.txt")
	gamma, epsilon := decode(input)
	return gamma * epsilon
}

func GetLifeSupportRating() uint64 {
	input := util.Read("binary_diagnostic/input.txt")
	o2 := decodeLifeSupportRatings(input, o2)
	co2 := decodeLifeSupportRatings(input, co2)
	return o2 * co2
}
