package seven_segment

import (
	"math"
	"math/bits"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

var m = map[rune]uint8{
	'a': 0, 'b': 1, 'c': 2,
	'd': 3, 'e': 4, 'f': 5,
	'g': 6}

func bitPattern(pattern string) uint8 {
	p := uint8(0)
	for _, c := range pattern {
		p = p | (1 << m[c])
	}
	return p
}

func signalToDigit(signals []string) map[uint8]uint8 {
	transcoded := make([]uint8, 10)
	sixSegmentDigits := []uint8{}
	fiveSegmentDigits := []uint8{}

	not := func(p uint8) uint8 {
		return (^p) & (^uint8(0) >> 1)
	}

	for i := range signals {
		p := bitPattern(signals[i])
		switch bits.OnesCount(uint(p)) {
		case 2:
			transcoded[1] = p
		case 3:
			transcoded[7] = p
		case 4:
			transcoded[4] = p
		case 7:
			transcoded[8] = p
		case 6:
			sixSegmentDigits = append(sixSegmentDigits, p)
		case 5:
			fiveSegmentDigits = append(fiveSegmentDigits, p)
		}
	}
	for i := range sixSegmentDigits {
		if transcoded[4] == sixSegmentDigits[i]&transcoded[4] {
			transcoded[9] = sixSegmentDigits[i]
		}
		if not(transcoded[1]) == sixSegmentDigits[i]&(not(transcoded[1])) {
			transcoded[6] = sixSegmentDigits[i]
		}
		if transcoded[4]^(not(transcoded[7])) == sixSegmentDigits[i]&(transcoded[4]^(not(transcoded[7]))) {
			transcoded[0] = sixSegmentDigits[i]
		}
	}

	for i := range fiveSegmentDigits {
		if not(transcoded[4]) == fiveSegmentDigits[i]&not(transcoded[4]) {
			transcoded[2] = fiveSegmentDigits[i]
		}
		if not(transcoded[4])^not(transcoded[1]) == fiveSegmentDigits[i]&(not(transcoded[4])^not(transcoded[1])) {
			transcoded[5] = fiveSegmentDigits[i]
		}
		if transcoded[1] == fiveSegmentDigits[i]&transcoded[1] {
			transcoded[3] = fiveSegmentDigits[i]
		}
	}

	m := make(map[uint8]uint8, 10)
	for k, v := range transcoded {
		m[v] = uint8(k)
	}
	return m
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Decode() int {
	data := util.Read("seven_segment/input.txt")
	sum := 0
	for _, line := range data {
		s := strings.Split(line, "|")
		signals := strings.Split(strings.TrimSpace(s[0]), " ")
		m := signalToDigit(signals)

		output := strings.Split(strings.TrimSpace(s[1]), " ")
		num := 0
		for i, o := range output {
			decoded := m[bitPattern(o)]
			num += int(decoded) * powInt(10, len(output)-1-i)
		}
		sum += num
	}
	return sum
}
