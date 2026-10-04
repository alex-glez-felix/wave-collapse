package wfc

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
