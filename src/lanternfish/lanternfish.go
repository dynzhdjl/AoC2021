package lanternfish

import (
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

func Population(days, spawn, offset int) int {
	return population(days, spawn, offset, "lanternfish/input.txt")
}

func population(days, spawn, offset int, file string) int {
	ledger := make([]int, spawn+offset+1)
	for _, v := range strings.Split(util.Read(file)[0], ",") {
		timer, _ := strconv.Atoi(v)
		ledger[timer]++
	}
	sum := func() int {
		cnt := 0
		for _, v := range ledger {
			cnt += v
		}
		return cnt
	}
	for i := 0; i < days; i++ {
		tmp := ledger[0]
		for i := 1; i < len(ledger); i++ {
			ledger[i-1] = ledger[i]
		}
		ledger[spawn] += tmp
		ledger[len(ledger)-1] = tmp
	}
	return sum()
}
