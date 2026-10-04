package wfc

import (
	"math/bits"
	"math/rand/v2"
)

type Cell uint64

func (c Cell) Has(t Tile) bool {
	return c&t.Bit() > 0
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
			allowed |= tile.Allowed(d)
		}
	}
	return allowed
}

func (c Cell) Collapse(r *rand.Rand) Cell {

	if c == 0 {
		panic("wfc: Collapse called on empty cell")
	}

	pick := r.IntN(c.Count())
	for tile := range NumTiles {
		if !c.Has(tile) {
			continue
		}
		if pick == 0 {
			return CellOf(tile)
		}
		pick--
	}
	return 0

}

func CellOf(tiles ...Tile) Cell {
	var cell Cell
	for _, tile := range tiles {
		cell |= tile.Bit()
	}
	return cell
}
