package bingo

type Card struct {
	Grid   [][]int
	Marked map[[2]int]int
}

func (c *Card) MarkSlot(slot [2]int) (bool, int) {
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
