package wfc

import "math/bits"

type Cell uint64

func (c Cell) Has(t Tile) bool {
	a := c & t.Bit()
	return a > 0
}

func (c Cell) Count() int {
	return bits.OnesCount64(uint64(c))
}

func (c Cell) Without(t Tile) Cell {
	return c &^ t.Bit()
}
