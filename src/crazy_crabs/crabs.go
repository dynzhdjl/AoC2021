package crazy_crabs

import (
	"sort"
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

func initilize(file string) ([]int, int, int) {
	arr := []int{}
	min := 0
	max := 0
	for _, v := range strings.Split(util.Read(file)[0], ",") {
		p, _ := strconv.Atoi(v)
		arr = append(arr, p)
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
	}
	return arr, min, max
}

func align(positions []int, min, max int) int {
	sort.Ints(positions)
	return (positions[49] + positions[50]) / 2
}

func HorizontalAlignmentPosition() int {
	return align(initilize("crazy_crabs/input.txt"))
}
