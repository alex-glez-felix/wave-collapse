package wfc

type Tile uint8

const (
	Water Tile = iota
	Sand
	Grass
	NumTiles
)

func (t Tile) Bit() Cell {
	return 1 << t

}
