package dumbo_octopus

import (
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
)

func countFlashes(octopus [][]int, steps int) int {
	n := len(octopus)
	queue := [][2]int{}
	flashCount := 0

	for s := 0; s < steps; s++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				octopus[i][j]++
				if octopus[i][j] >= 10 {
					queue = append(queue, [2]int{i, j})
				}
			}
		}
		for len(queue) > 0 {
			position := queue[0]
			queue = queue[1:]
			flashCount += flash(octopus, position)
		}
	}
	return flashCount
}

func findSyncronizationStep(octopus [][]int) int {
	n := len(octopus)
	queue := [][2]int{}
	steps := 0
	for {
		flashCount := 0
		steps++
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				octopus[i][j]++
				if octopus[i][j] >= 10 {
					queue = append(queue, [2]int{i, j})
				}
			}
		}
		for len(queue) > 0 {
			position := queue[0]
			queue = queue[1:]
			flashCount += flash(octopus, position)
		}
		if flashCount == n*n {
			break
		}

	}
	return steps
}

func flash(octopus [][]int, position [2]int) int {
	dirs := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	n := len(octopus)
	i := position[0]
	j := position[1]
	sum := 0
	octopus[i][j] = 0
	for _, dir := range dirs {
		di := i + dir[0]
		dj := j + dir[1]
		if di < 0 || dj < 0 || di >= n || dj >= n {
			continue
		}
		if octopus[di][dj] == 0 {
			continue
		}
		octopus[di][dj]++
		if octopus[di][dj] == 10 {
			sum += flash(octopus, [2]int{di, dj})
		}
	}
	return 1 + sum
}

func CountFlashes(steps int) int {
	octopus := getOctopusMap("dumbo_octopus/input.txt")
	return countFlashes(octopus, steps)
}

func FindSyncronizationStep() int {
	octopus := getOctopusMap("dumbo_octopus/input.txt")
	return findSyncronizationStep(octopus)
}

func getOctopusMap(file string) [][]int {
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
