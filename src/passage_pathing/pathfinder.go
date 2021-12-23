package passage_pathing

import (
	"strings"
	"unicode"

	"github.com/dynzhdjl/AoC2021/util"
)

func getNumberOfPaths(graph map[string][]string, from, to string, visits map[string]bool) int {
	cnt := 0
	var walk func(string)
	walk = func(current string) {
		if current == to {
			cnt++
			return
		}
		if visited, _ := visits[current]; visited && !startsWithUpper(current) {
			return
		}
		visits[current] = true
		adjacents, _ := graph[current]
		for _, adjacent := range adjacents {
			walk(adjacent)
		}
		visits[current] = false
	}
	walk(from)
	return cnt
}

func createAdjancenyList(input []string) map[string][]string {
	g := map[string][]string{}
	for _, edge := range input {
		verticies := strings.Split(edge, "-")
		_, ok := g[verticies[0]]
		if !ok {
			g[verticies[0]] = []string{verticies[1]}
		} else {
			g[verticies[0]] = append(g[verticies[0]], verticies[1])
		}
		_, ok = g[verticies[1]]
		if !ok {
			g[verticies[1]] = []string{verticies[0]}
		} else {
			g[verticies[1]] = append(g[verticies[1]], verticies[0])
		}
	}
	return g
}

func startsWithUpper(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

func PathCount() int {
	intput := util.Read("passage_pathing/input.txt")
	g := createAdjancenyList(intput)
	return getNumberOfPaths(g, "start", "end", map[string]bool{})
}
