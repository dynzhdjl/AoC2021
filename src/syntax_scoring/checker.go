package syntax_scoring

import (
	"errors"
	"fmt"

	"github.com/dynzhdjl/AoC2021/util"
)

var illigalSymbolScoreMap = map[byte]int{
	'}': 1197,
	')': 3,
	'>': 25137,
	']': 57}

var autoCompleteSymbolScoreMap = map[byte]int{
	'}': 3,
	')': 1,
	'>': 4,
	']': 2}

var openCloseSymbols = map[byte]byte{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')'}

func findIlligalSymbol(line string) (int, error, []byte) {
	opens := []byte{}
	closing := func(symbol byte) bool {
		return symbol == '}' || symbol == ']' || symbol == '>' || symbol == ')'
	}

	for i := 0; i < len(line); i++ {
		if !closing(line[i]) {
			opens = append(opens, line[i])
		} else {
			if len(opens) == 0 {
				return i,
					errors.New(fmt.Sprintf("Unexpected symbol %c found! At index: %d", line[i], i)), opens
			}
			expected, _ := openCloseSymbols[opens[len(opens)-1]]
			opens = opens[:len(opens)-1]
			found := line[i]
			if found != expected {
				return i, errors.New(fmt.Sprintf("Expected %c, but found %c. At index: %d", expected, found, i)), opens
			}
		}
	}
	return len(line), nil, opens
}

func CalculateSyntaxCheckerScore() int {
	input := util.Read("syntax_scoring/input.txt")
	totalScore := 0
	for _, v := range input {
		if i, err, _ := findIlligalSymbol(v); err != nil {
			symbolScore, _ := illigalSymbolScoreMap[v[i]]
			totalScore += symbolScore
		}
	}
	return totalScore
}

func CalculateAutoCompleteScore() int {
	input := util.Read("syntax_scoring/input.txt")

	scores := []int{}
	for _, v := range input {
		totalScore := 0
		if _, err, opens := findIlligalSymbol(v); err == nil && len(opens) > 0 {
			for i := len(opens) - 1; i >= 0; i-- {
				closing, _ := openCloseSymbols[opens[i]]
				score, _ := autoCompleteSymbolScoreMap[closing]
				totalScore = (totalScore * 5) + score
			}
			scores = append(scores, totalScore)
		}
	}
	midScore := 0
	heap := util.MinHeap(scores)
	mid := len(scores) / 2
	for i := 0; i <= mid; i++ {
		midScore = util.PeekHeap(&heap)
	}
	return midScore
}
