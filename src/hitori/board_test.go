package hitori

import "testing"

func TestSolvePuzzle(t *testing.T) {
	cellData := [][]int{
		[]int{298386517},
		[]int{421461934},
		[]int{453638721},
		[]int{385215699},
		[]int{361517914},
		[]int{952253776},
		[]int{335987462},
		[]int{816295244},
		[]int{473129658},
	}

	board := PopulateBoard(cellData)
	board.Solve()
    if !board.isSolved() {
        t.Error("Puzzle not solved")
    }
}
