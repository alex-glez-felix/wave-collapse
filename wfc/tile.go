package wfc

import "fmt"

type Tile uint8

const (
	Water Tile = iota
	Sand
	Grass
	NumTiles
)

const AllTiles = 1<<NumTiles - 1

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
	NumDirections
)

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

var compatible = [NumTiles]Cell{
	Water: CellOf(Water, Sand),
	Sand:  AllTiles,
	Grass: CellOf(Grass, Sand),
}

func (t Tile) Bit() Cell {
	return 1 << t
}

func (t Tile) Allowed(d Direction) Cell {
	return compatible[t]
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
