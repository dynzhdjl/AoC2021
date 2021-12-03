package binary_diagnostic

import (
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
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

func GetEnergyConsumption() uint64 {
	input := util.Read("binary_diagnostic/input.txt")
	gamma, epsilon := decode(input)
	return gamma * epsilon
}
