package hitori

type cell struct {
	value  int
	black  bool
	yellow bool
}

func unknownCell(value int) *cell {
	c := new(cell)
	c.value = value
	c.black = false
	c.yellow = false
	return c
}

func (this *cell) setBlack() bool {
	if this.yellow {
		panic("Cell is already yellow, cannot make black")
	}

	wasBlack := this.black
	this.black = true
	return !wasBlack
}

func (this *cell) setYellow() bool {
	if this.black {
		panic("Cell is already black, cannot make yellow")
	}

	wasYellow := this.yellow
	this.yellow = true
	return !wasYellow
}

func (this *cell) isYellow() bool {
	return this.yellow
}

func (this *cell) isBlack() bool {
	return this.black
}

func (this *cell) isUnknown() bool {
	return !this.black && !this.yellow
}

func (this *cell) getValue() int {
	return this.value
}
