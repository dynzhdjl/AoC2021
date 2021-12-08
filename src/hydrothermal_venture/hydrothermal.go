package hydrothermal_venture

import (
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

type Point struct {
	x int
	y int
}

func newPoint(input string) Point {
	points := strings.Split(input, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return Point{x, y}
}

type Line struct {
	start     Point
	end       Point
	direction [2]int
}

func newLine(start, end string) Line {
	s := newPoint(start)
	e := newPoint(end)
	dx, dy := getDirecton(e.x, s.x), getDirecton(e.y, s.y)
	return Line{s, e, [2]int{dx, dy}}
}

func getDirecton(end, start int) int {
	r := end - start
	if r == 0 {
		return 0
	}
	return r / abs(r)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -1 * a
}

func parseRawData(input []string) ([]Line, int, int) {
	m, n := 0, 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	line := make([]Line, len(input))
	for k, v := range input {
		pair := strings.Split(v, "->")
		start := strings.TrimSpace(pair[0])
		end := strings.TrimSpace(pair[1])
		line[k] = newLine(start, end)
		m = max(m, max(line[k].start.x, line[k].end.x))
		n = max(n, max(line[k].start.y, line[k].end.y))
	}
	return line, m + 1, n + 1
}

func NumberOfVentLineIntersection(minIntersection int) int {
	lines, m, n := parseRawData(util.Read("hydrothermal_venture/input.txt"))
	terrain := make([][]int, n)
	for i := range terrain {
		terrain[i] = make([]int, m)
	}
	walk := func(l Line) int {
		x := l.start.x
		y := l.start.y
		cnt := 0
		for {
			terrain[y][x]++
			if terrain[y][x] == minIntersection {
				cnt++
			}
			if [2]int{l.end.x, l.end.y} == [2]int{x, y} {
				break
			}
			x = x + l.direction[0]
			y = y + l.direction[1]
		}
		return cnt
	}
	intersectionCnt := 0
	for _, l := range lines {
		intersectionCnt += walk(l)
	}
	return intersectionCnt
}
