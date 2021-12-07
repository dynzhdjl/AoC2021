package bingo

import (
	"testing"
)

func TestBingoIsWinner(t *testing.T) {
	//nums := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	cards := [][5][5]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}

	//7,4,9,5,11,17,23,2,0,14,21,24

	card := Card{cards[2], map[[2]int]int{}, false}
	winner, score := card.markSlot([2]int{4, 4})
	expectedWinner := false
	expectedScore := -1

	if winner != expectedWinner {
		t.Errorf("expected: %v, got: %v", expectedWinner, winner)
	}
	if score != expectedScore {
		t.Errorf("expected: %v, got: %v", expectedScore, score)
	}

	card.markSlot([2]int{0, 4})
	card.markSlot([2]int{1, 3})
	card.markSlot([2]int{3, 4})
	card.markSlot([2]int{3, 1})
	card.markSlot([2]int{0, 2})
	card.markSlot([2]int{2, 2})
	card.markSlot([2]int{4, 0})
	card.markSlot([2]int{4, 1})
	card.markSlot([2]int{0, 0})
	card.markSlot([2]int{0, 1})

	expectedWinner = true
	expectedScore = 4512
	winner, score = card.markSlot([2]int{0, 3})

	if winner != expectedWinner {
		t.Errorf("expected: %v, got: %v", expectedWinner, winner)
	}
	if score != expectedScore {
		t.Errorf("expected: %v, got: %v", expectedScore, score)
	}
}
