package extended_polymerization

import (
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

func Diff(n int) uint64 {
	template, rules := parseRawData("extended_polymerization/input.txt")
	elements := build(template, rules, n)
	min, max := minMax(elements)
	minCnt, _ := elements[min]
	maxCnt, _ := elements[max]
	return maxCnt - minCnt
}

func build(template string, rules map[string]byte, n int) map[byte]uint64 {
	stats := map[string]uint64{}
	elementCount := map[byte]uint64{}
	var e1 byte
	var e2 byte
	for i := 1; i < len(template); i++ {
		e1 = template[i-1]
		e2 = template[i]
		elementCount[e1]++
		stats[string(e1)+string(e2)]++
	}
	elementCount[e2]++
	for i := 0; i < n; i++ {
		tmp := map[string]uint64{}
		for k, v := range stats {
			element, _ := rules[k]
			elementCount[element] += v
			tmp[string(k[0])+string(element)] += v
			tmp[string(element)+string(k[1])] += v
		}
		stats = tmp
	}
	return elementCount
}

func minMax(elements map[byte]uint64) (byte, byte) {
	min := ^uint64(0)
	max := uint64(0)
	var minElement byte
	var maxElement byte
	for e, c := range elements {
		if c <= min {
			minElement = e
			min = c
		}
		if c >= max {
			maxElement = e
			max = c
		}
	}
	return minElement, maxElement
}

func parseRawData(file string) (string, map[string]byte) {
	lines := util.Read(file)
	template := lines[0]
	rules := makeRules(lines[2:])
	return template, rules
}

func makeRules(lines []string) map[string]byte {
	r := map[string]byte{}
	for _, l := range lines {
		p := strings.Split(l, "->")
		r[strings.TrimSpace(p[0])] = byte(strings.TrimSpace(p[1])[0])
	}
	return r
}
