package hitori

import "strconv"
import "fmt"

type board struct {
	cells [][]*cell
}

func PopulateBoard(cells [][]int) *board {
	b := new(board)
	b.cells = make([][]*cell, len(cells))
	for i := 0; i < len(cells); i++ {
		b.cells[i] = make([]*cell, len(cells[i]))
		for j := 0; j < len(cells[i]); j++ {
			b.cells[i][j] = unknownCell(cells[i][j])
		}
	}

	return b
}

func (this *board) getCellValues() [][]int {
	values := make([][]int, len(this.cells))
	for i := 0; i < len(this.cells); i++ {
		values[i] = make([]int, len(this.cells[i]))
		for j := 0; j < len(this.cells[i]); j++ {
			values[i][j] = this.cells[i][j].getValue()
		}
	}

	return values
}

func (this *board) clone() *board {
	clonedBoard := PopulateBoard(this.getCellValues())

	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if this.cells[i][j].isBlack() {
				clonedBoard.cells[i][j].setBlack()
			} else if this.cells[i][j].isYellow() {
				clonedBoard.cells[i][j].setYellow()
			}
		}
	}

	return clonedBoard
}

func (this *board) String() string {
	output := ""
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			var surround = " "
			if this.cells[i][j].isYellow() {
				surround = " "
			}

			if this.cells[i][j].isBlack() {
				surround = "*"
			}

			output += " " + surround + strconv.Itoa(this.cells[i][j].getValue()) + surround + " "
		}

		output += "\n"
	}

	return output
}

func (this *board) Solve() {
	var i int
	for i = 0; i < 20 && !this.isSolved(); i++ {
		this.markNonDuplicates()
		this.tryBlacks()
		this.tryYellows()
	}

	fmt.Println(i, "Rounds")
}

func (this *board) markNonDuplicates() {
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if !this.findDuplicates(i, j) {
				this.cells[i][j].setYellow()
			}
		}
	}
}

func (this *board) findDuplicates(x, y int) bool {
	for i := 0; i < len(this.cells); i++ {
		if i != x && this.cells[i][y].getValue() == this.cells[x][y].getValue() && !this.cells[i][y].isBlack() {
			return true
		}
	}

	for j := 0; j < len(this.cells[x]); j++ {
		if j != y && this.cells[x][j].getValue() == this.cells[x][y].getValue() && !this.cells[x][j].isBlack() {
			return true
		}
	}

	return false
}

func (this *board) tryBlacks() {
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if this.cells[i][j].isUnknown() {
				b := this.clone()
				if !b.tryBlack(i, j) {
					this.cells[i][j].setYellow()
				}
			}
		}
	}
}

func (this *board) tryYellows() {
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if this.cells[i][j].isUnknown() {
				b := this.clone()
				if !b.tryYellow(i, j) {
					this.cells[i][j].setBlack()
				}
			}
		}
	}
}

func (this *board) tryBlack(x, y int) bool {
	if !this.setBlackAndVerify(x, y) {
		return false
	}

	return true
}

func (this *board) setBlackAndVerify(x, y int) bool {
	this.cells[x][y].setBlack()
	if (x > 0 && this.cells[x-1][y].isBlack()) ||
		(y > 0 && this.cells[x][y-1].isBlack()) ||
		(x < len(this.cells)-1 && this.cells[x+1][y].isBlack()) ||
		(y < len(this.cells[x])-1 && this.cells[x][y+1].isBlack()) {
		return false
	}

	if !this.connected() {
		return false
	}

	return this.setBlackConclusions(x, y)
}

func (this *board) tryYellow(x, y int) bool {
	if !this.setYellowAndVerify(x, y) {
		return false
	}

	return true
}

func (this *board) setYellowAndVerify(x, y int) bool {
	this.cells[x][y].setYellow()
	for i := 0; i < len(this.cells); i++ {
		if i != x &&
			this.cells[i][y].getValue() == this.cells[x][y].getValue() &&
			this.cells[i][y].isYellow() {
			return false
		}
	}

	for j := 0; j < len(this.cells[x]); j++ {
		if j != y &&
			this.cells[x][j].getValue() == this.cells[x][y].getValue() &&
			this.cells[x][j].isYellow() {
			return false
		}
	}

	return this.setYellowConclusions(x, y)
}

func (this *board) setYellowConclusions(x, y int) bool {
	for i := 0; i < len(this.cells); i++ {
		if i != x &&
			this.cells[i][y].getValue() == this.cells[x][y].getValue() &&
			this.cells[i][y].isUnknown() &&
			!this.setBlackAndVerify(i, y) {
			return false
		}
	}

	for j := 0; j < len(this.cells[x]); j++ {
		if j != y &&
			this.cells[x][j].getValue() == this.cells[x][y].getValue() &&
			this.cells[x][j].isUnknown() &&
			!this.setBlackAndVerify(x, j) {
			return false
		}
	}

	return true
}

func (this *board) setBlackConclusions(x, y int) bool {
	if x > 0 && !this.setYellowAndVerify(x-1, y) {
		return false
	}

	if y > 0 && !this.setYellowAndVerify(x, y-1) {
		return false
	}

	if x < len(this.cells)-1 && !this.setYellowAndVerify(x+1, y) {
		return false
	}

	if y < len(this.cells[x])-1 && !this.setYellowAndVerify(x, y+1) {
		return false
	}

	return true
}

func (this *board) isSolved() bool {
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if this.cells[i][j].isUnknown() {
				return false
			}
		}
	}

	return true
}

func (this *board) connected() bool {
	var (
		nonBlacks []*cell
		x         int
		y         int
	)
	for i := 0; i < len(this.cells); i++ {
		for j := 0; j < len(this.cells[i]); j++ {
			if !this.cells[i][j].isBlack() {
				nonBlacks = append(nonBlacks, this.cells[i][j])
				x = i
				y = j
			}
		}
	}

	connections := this.getConnections(x, y, make([]*cell, 0))

	return len(connections) == len(nonBlacks)
}

func (this *board) getConnections(x, y int, found []*cell) []*cell {
	found = append(found, this.cells[x][y])
	if x > 0 &&
		!contains(this.cells[x-1][y], found) &&
		!this.cells[x-1][y].isBlack() {
		newFound := this.getConnections(x-1, y, found)
		for _, val := range newFound {
			if !contains(val, found) {
				found = append(found, val)
			}
		}
	}

	if x < len(this.cells)-1 &&
		!contains(this.cells[x+1][y], found) &&
		!this.cells[x+1][y].isBlack() {
		newFound := this.getConnections(x+1, y, found)
		for _, val := range newFound {
			if !contains(val, found) {
				found = append(found, val)
			}
		}
	}

	if y > 0 &&
		!contains(this.cells[x][y-1], found) &&
		!this.cells[x][y-1].isBlack() {
		newFound := this.getConnections(x, y-1, found)
		for _, val := range newFound {
			if !contains(val, found) {
				found = append(found, val)
			}
		}
	}

	if y < len(this.cells[x])-1 &&
		!contains(this.cells[x][y+1], found) &&
		!this.cells[x][y+1].isBlack() {
		newFound := this.getConnections(x, y+1, found)
		for _, val := range newFound {
			if !contains(val, found) {
				found = append(found, val)
			}
		}
	}

	return found
}

func contains(needle *cell, haystack []*cell) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}
