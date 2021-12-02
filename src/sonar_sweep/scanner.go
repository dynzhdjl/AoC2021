package sonar_sweep

import (
	"fmt"
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
)

func scanner(input []int) int {
	n := len(input)
	cnt := 0
	p := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			p = input[i]
			continue
		}
		c := input[i]
		if c > p {
			cnt++
		}
		p = c
	}
	return cnt
}

func slidingScanner(input []int) int {
	n := len(input)
	cnt := 0
	p := 0
	for i := 0; i < n-2; i++ {
		if i == 0 {
			p = input[i] + input[i+1] + input[i+2]
			continue
		}
		c := p - input[i-1] + input[i+2]
		if c > p {
			cnt++
		}
		p = c
	}
	return cnt
}

func sweep(scanner func([]int) int) int {
	lines := util.Read("sonar_sweep/input.txt")
	input := make([]int, len(lines))
	for _, l := range lines {
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(fmt.Sprintf("could parse the input: %s", l))
		}
		input = append(input, v)
	}
	return scanner(input)
}

func Sweep() int {
	return sweep(scanner)
}

func SlidingSweep() int {
	return sweep(slidingScanner)
}
