package smoke_basin

import (
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
)

func lowPoints(m [][]int) [][2]int {
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	rows := len(m)
	cols := len(m[0])
	lows := [][2]int{}

	low := func(row, col int) bool {
		isLow := true
		for _, d := range dirs {
			dr := row + d[0]
			dc := col + d[1]
			if dc < 0 || dr < 0 || dr >= rows || dc >= cols {
				continue
			}
			if m[row][col] < m[dr][dc] {
				continue
			}
			isLow = false
		}
		return isLow
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if low(row, col) {
				lows = append(lows, [2]int{row, col})
			}
		}
	}
	return lows
}

func explore(hightmap [][]int, row, col int) int {
	if row < 0 || row >= len(hightmap) || col < 0 || col >= len(hightmap[0]) {
		return 0
	}
	if hightmap[row][col] >= 9 || hightmap[row][col] < 0 {
		return 0
	}
	hightmap[row][col] = -1
	return 1 +
		explore(hightmap, row+1, col) +
		explore(hightmap, row-1, col) +
		explore(hightmap, row, col+1) +
		explore(hightmap, row, col-1)
}

func getHightMap(file string) [][]int {
	rows := util.Read(file)
	m := [][]int{}
	for _, row := range rows {
		r := []int{}
		for _, n := range row {
			n, _ := strconv.Atoi(string(n))
			r = append(r, n)
		}
		m = append(m, r)
	}
	return m
}

func FindBiggestBasin() int {
	m := getHightMap("smoke_basin/input.txt")

	largestAreas := []int{0}
	for _, v := range lowPoints(m) {
		area := explore(m, v[0], v[1])
		if area > largestAreas[len(largestAreas)-1] {
			if len(largestAreas) >= 3 {
				largestAreas = largestAreas[1:]
			}
			largestAreas = append(largestAreas, area)
		}
	}
	return largestAreas[0] * largestAreas[1] * largestAreas[2]
}

func CalculateRisk() int {
	m := getHightMap("smoke_basin/input.txt")
	risk := 0
	for _, v := range lowPoints(m) {
		risk += (m[v[0]][v[1]] + 1)
	}
	return risk
}
