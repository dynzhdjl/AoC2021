package bingo

import (
	"strconv"
	"strings"

	"github.com/dynzhdjl/AoC2021/util"
)

type Card struct {
	Grid   [5][5]int
	Marked map[[2]int]int
}

func (c *Card) markSlot(slot [2]int) (bool, int) {
	c.Marked[slot] = c.Grid[slot[0]][slot[1]]
	isWinner := c.isWinner(slot)
	if isWinner {
		return isWinner, c.getScore(slot)
	}
	return isWinner, -1
}

func (c *Card) getScore(slot [2]int) int {
	score := 0
	n := len(c.Grid)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if _, ok := c.Marked[[2]int{row, col}]; !ok {
				score += c.Grid[row][col]
			}
		}
	}
	return score * c.Grid[slot[0]][slot[1]]
}

func (c *Card) isWinner(slot [2]int) bool {
	n := len(c.Grid)
	if len(c.Marked) < n {
		return false
	}
	row := slot[0]
	col := slot[1]
	isWinner := true
	for i := 0; i < n; i++ {
		if _, ok := c.Marked[[2]int{row, i}]; !ok {
			isWinner = false
		}
	}
	if isWinner {
		return isWinner
	}
	for i := 0; i < n; i++ {
		if _, ok := c.Marked[[2]int{i, col}]; !ok {
			isWinner = false
		}
	}
	return isWinner
}

type CardIndex struct {
	card *Card
	slot [2]int
}

var cardIndex map[int]*[]CardIndex = map[int]*[]CardIndex{}

func WinnerScore() int {
	input := util.Read("bingo/input.txt")
	draws := []int{}
	for _, v := range strings.Split(input[0], ",") {
		i, _ := strconv.Atoi(v)
		draws = append(draws, i)
	}

	for i := 2; i < len(input); {
		if input[i] != "\n" {
			c := &Card{[5][5]int{}, map[[2]int]int{}}
			for row := 0; row < 5; row++ {
				cols := strings.Fields(input[row+i])
				for j := 0; j < len(cols); j++ {
					v, _ := strconv.Atoi(cols[j])
					c.Grid[row][j] = v
					if index, ok := cardIndex[v]; ok {
						*index = append(*index, CardIndex{c, [2]int{row, j}})
					} else {
						cardIndex[v] = &[]CardIndex{{c, [2]int{row, j}}}
					}
				}
			}
		}
		i += 6
	}

	winner := false
	score := 0
	for i := 0; i < len(draws); i++ {
		if cards, ok := cardIndex[draws[i]]; ok {
			for _, c := range *cards {
				winner, score = c.card.markSlot(c.slot)
				if winner {
					return score
				}
			}
		}
	}
	return score
}
