package chiton

import (
	"container/heap"
	"strconv"

	"github.com/dynzhdjl/AoC2021/util"
)

type Item struct {
	risk     int
	position [2]int
}

type RiskHeap []Item

func (h RiskHeap) Len() int           { return len(h) }
func (h RiskHeap) Less(i, j int) bool { return h[i].risk < h[j].risk }
func (h RiskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *RiskHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *RiskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getRiskMap(file string) [][]int {
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

func calculator(riskMap [][]int) int {
	auxilary := make([][]int, len(riskMap))
	for i := range auxilary {
		auxilary[i] = make([]int, len(riskMap[0]))
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	rows := len(riskMap)
	cols := len(riskMap[0])
	pq := RiskHeap{}
	pq = append(pq, Item{riskMap[0][1], [2]int{0, 1}})
	pq = append(pq, Item{riskMap[1][0], [2]int{1, 0}})
	auxilary[0][0] *= -1
	heap.Init(&pq)
	for len(pq) > 0 {
		item := heap.Pop(&pq).(Item)
		for _, dir := range dirs {
			dy := item.position[0] + dir[0]
			dx := item.position[1] + dir[1]
			if dx < 0 || dy < 0 || dx >= cols || dy >= rows {
				continue
			}
			if auxilary[dy][dx] == 0 {
				auxilary[dy][dx] = item.risk + riskMap[dy][dx]
				heap.Push(&pq, Item{auxilary[dy][dx], [2]int{dy, dx}})
			}
		}
		auxilary[item.position[0]][item.position[1]] *= -1
	}
	return -1 * auxilary[rows-1][cols-1]
}

func Calculate() int {
	riskMap := getRiskMap("chiton/input.txt")
	return calculator(riskMap)
}

func Calculate5x() int {
	riskMap := getRiskMap("chiton/input.txt")
	n := len(riskMap)
	big := make([][]int, n*5)
	for i := range big {
		big[i] = make([]int, n*5)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					y := k + i*n
					x := l + j*n
					v := riskMap[k][l] + i + j
					if v > 9 {
						v = (v % 9)
					}
					big[y][x] = v
				}
			}
		}
	}
	return calculator(big)
}
