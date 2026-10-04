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

func (c Cell) Allowed(d Direction) Cell {

	var allowed Cell
	for tile := range NumTiles {
		if c.Has(tile) {
			allowed = allowed | tile.Allowed(d)
		}
	}
	return allowed
}

func CellOf(tiles ...Tile) Cell {
	var cell Cell
	for _, tile := range tiles {
		cell = cell | tile.Bit()
	}
	return cell
}
