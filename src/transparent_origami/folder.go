package transparent_origami

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

func foldAlongX(start, end [2]int, grid [][]int, p int) ([2]int, [2]int) {
	if p < start[1] || p > end[1] || p < 0 || p >= len(grid[0]) {
		panic("invalid argument")
	}
	s, e := start, end
	ns, ne := start, end
	f := 0
	if abs(p-start[1]) < abs(end[1]-p) {
		e[1] = p - 1
		ns[1] = p + 1
		f = 1
	} else {
		s[1] = p + 1
		ne[1] = p - 1
		f = -1
	}

	for y := s[0]; y <= e[0]; y++ {
		grid[y][p] = 0
		for x := s[1]; x <= e[1]; x++ {
			if grid[y][x] == 1 {
				dx := p + (f * abs(p-x))
				grid[y][dx] = grid[y][x]
				grid[y][x] = 0
			}
		}
	}
	return ns, ne
}

func foldAlongY(start, end [2]int, grid [][]int, p int) ([2]int, [2]int) {
	if p < start[0] || p > end[0] || p < 0 || p >= len(grid) {
		panic("invalid argument")
	}
	fs, fe := start, end
	ns, ne := start, end
	f := 0
	if abs(p-start[0]) < abs(end[0]-p) {
		fe[0] = p - 1
		ns[0] = p + 1
		f = 1
	} else {
		fs[0] = p + 1
		ne[0] = p - 1
		f = -1
	}
	for x := fs[1]; x <= fe[1]; x++ {
		grid[p][x] = 0
		for y := fs[0]; y <= fe[0]; y++ {
			if grid[y][x] == 1 {
				dy := p + (f * abs(p-y))
				grid[dy][x] = grid[y][x]
				grid[y][x] = 0
			}
		}
	}
	return ns, ne
}

func count(start, end [2]int, grid Paper) int {
	cnt := 0
	for y := start[0]; y <= end[0]; y++ {
		for x := start[1]; x <= end[1]; x++ {
			if grid[y][x] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

type Paper [][]int

func (p Paper) applyInstruction(start, end [2]int, i FoldingInsturction) ([2]int, [2]int) {
	var ns, ne [2]int
	switch i.axis {
	case "x":
		ns, ne = foldAlongX(start, end, p, i.val)
	case "y":
		ns, ne = foldAlongY(start, end, p, i.val)
	default:
		panic("unknown instruction")
	}
	return ns, ne
}

func Count() int {
	dots, m, n, instuctions := parseRawData(util.Read("transparent_origami/input.txt"))
	var paper Paper
	paper = make([][]int, m)
	for i := range paper {
		paper[i] = make([]int, n)
	}

	for _, dot := range dots {
		paper[dot[0]][dot[1]] = 1
	}

	start, end := [2]int{0, 0}, [2]int{len(paper) - 1, len(paper[0]) - 1}
	for _, instruction := range instuctions {
		start, end = paper.applyInstruction(start, end, instruction)
	}
	for y := start[0]; y <= end[0]; y++ {
		for x := start[1]; x <= end[1]; x++ {
			if paper[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	return count(start, end, paper)
}

type FoldingInsturction struct {
	axis string
	val  int
}

func parseRawData(input []string) ([][2]int, int, int, []FoldingInsturction) {
	m, n := 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dots := [][2]int{}
	index := 0
	for k, v := range input {
		if v == "========" {
			index = k
			break
		}
		pair := strings.Split(v, ",")
		y, _ := strconv.Atoi(pair[1])
		x, _ := strconv.Atoi(pair[0])
		dots = append(dots, [2]int{y, x})
		m = max(m, y)
		n = max(n, x)
	}

	input = input[index+1:]
	instructions := []FoldingInsturction{}
	for _, v := range input {
		s := strings.Split(strings.Split(v, " ")[2], "=")
		axis := s[0]
		val, _ := strconv.Atoi(s[1])
		instructions = append(instructions, FoldingInsturction{axis: axis, val: val})
	}
	return dots, m + 1, n + 1, instructions
}
