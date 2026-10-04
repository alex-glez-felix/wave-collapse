package wfc

import "fmt"

type Tile uint8

const (
	Water Tile = iota
	Sand
	Grass
	NumTiles
)

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
}

type TileSet struct {
	tiles []tileInfo
}

func (ts *TileSet) Add(name string, glyph rune) Tile {
	if len(ts.tiles) >= 64 {
		panic("wfc: Cannot add more than 64 tiles to a single tileset")
	}

	tile := Tile(len(ts.tiles))
	ts.tiles = append(ts.tiles, tileInfo{name: name, glyph: glyph})
	return tile
}

func (ts *TileSet) Connect(t1, t2 Tile) {
	for d := range NumDirections {
		ts.tiles[t1].rules[d] |= t2.Bit()
		ts.tiles[t2].rules[d] |= t1.Bit()
	}
}

const AllTiles = 1<<NumTiles - 1

var rules = [NumTiles][NumDirections]Cell{
	Water: {
		Up:    CellOf(Water, Sand),
		Right: CellOf(Water),
		Down:  CellOf(Water, Sand),
		Left:  CellOf(Water, Sand),
	},
	Sand: {
		Up:    CellOf(Water, Sand, Grass),
		Right: CellOf(Sand, Water),
		Down:  CellOf(Water, Sand, Grass),
		Left:  CellOf(Grass, Sand),
	},
	Grass: {
		Up:    CellOf(Grass, Sand),
		Right: CellOf(Grass, Sand),
		Down:  CellOf(Grass, Sand),
		Left:  CellOf(Grass),
	},
}

func (t Tile) Bit() Cell {
	return 1 << t
}

func (t Tile) Allowed(d Direction) Cell {
	return rules[t][d]
}

func (t Tile) String() string {
	switch t {
	case Water:
		return "Water"
	case Sand:
		return "Sand"
	case Grass:
		return "Grass"
	default:
		return fmt.Sprintf("Tile(%d)", t)
	}
}

func (t Tile) Glyph() rune {
	switch t {
	case Water:
		return '.'
	case Sand:
		return '~'
	case Grass:
		return '#'
	default:
		return '?'
	}
}
