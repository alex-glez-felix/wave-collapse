package wfc

type Tile uint8
type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
	NumDirections
)

func (d Direction) Opposite() Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	}

	panic("wfc: Direction should always have an opposite")
}

type tileInfo struct {
	name  string
	glyph rune
	rules [NumDirections]Cell
	color uint8
}

type TileSet struct {
	tiles []tileInfo
}

func (ts *TileSet) Add(name string, glyph rune, color uint8) Tile {
	if len(ts.tiles) >= 64 {
		panic("wfc: Cannot add more than 64 tiles to a single tileset")
	}

	tile := Tile(len(ts.tiles))
	ts.tiles = append(ts.tiles, tileInfo{name: name, glyph: glyph, color: color})
	return tile
}

func (ts *TileSet) Connect(t1, t2 Tile) {
	for d := range NumDirections {
		ts.tiles[t1].rules[d] |= t2.Bit()
		ts.tiles[t2].rules[d] |= t1.Bit()
	}
}

func (ts *TileSet) All() Cell {
	return 1<<len(ts.tiles) - 1
}

func (ts *TileSet) Allowed(t Tile, d Direction) Cell {
	return ts.tiles[t].rules[d]
}

func (ts *TileSet) AllowedFrom(c Cell, d Direction) Cell {
	var allowed Cell
	for i := range ts.tiles {
		if c.Has(Tile(i)) {
			allowed |= ts.tiles[i].rules[d]
		}
	}
	return allowed
}

func (t Tile) Bit() Cell {
	return 1 << t
}
