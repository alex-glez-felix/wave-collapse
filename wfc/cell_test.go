package wfc

import "testing"

func TestCellHas(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
		tile Tile
		want bool
	}{
		{"Water present", 0b0001, Water, true},
		{"Water not present", 0b0000, Water, false},
		{"Water absent others present", 0b1110, Water, false},
		{"Sand present", 0b0011, Sand, true},
		{"Grass present", 0b0111, Grass, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Has(tt.tile); got != tt.want {
				t.Errorf("Has(%v) = %v, want %v", tt.tile, got, tt.want)
			}
		})
	}
}

func TestCellCount(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
		want int
	}{
		{"Four ones", 0b1111, 4},
		{"Empty", 0b0000, 0},
		{"2 Ones", 0b0011, 2},
		{"3 ones", 0b1101, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Count(); got != tt.want {
				t.Errorf("Count(%v) = %v, want %v", tt.cell, got, tt.want)
			}
		})
	}
}

func TestCellWithout(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
		tile Tile
		want Cell
	}{
		{"Remove water", 0b0001, Water, 0b0000},
		{"Remove water from nothing", 0b0000, Water, 0b0000},
		{"Remove sand", 0b1111, Sand, 0b1101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Without(tt.tile); got != tt.want {
				t.Errorf("Without(%v) on %v = %v, want %v", tt.tile, tt.cell, got, tt.want)
			}
		})
	}
}

func TestCellOf(t *testing.T) {
	tests := []struct {
		name string
		of   []Tile
		want Cell
	}{
		{"Empty cell", []Tile{}, 0b000},
		{"Water cell", []Tile{Water}, 0b001},
		{"Water and Sand", []Tile{Water, Sand}, 0b011},
		{"Sand and Grass", []Tile{Grass, Sand}, 0b110},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CellOf(tt.of...); got != tt.want {
				t.Errorf("CellOf(%v) = %v, want %v", tt.of, got, tt.want)
			}
		})
	}
}

func TestCellAllowed(t *testing.T) {
	tests := []struct {
		name      string
		cell      Cell
		direction Direction
		want      Cell
	}{
		{"Cell Water Up allowed", CellOf(Water), Up, CellOf(Water, Sand)},
		{"Cell Water, Sand Up allowed", CellOf(Water, Sand), Up, CellOf(Water, Sand, Grass)},
		{"Cell Water, Sand, Grass Up allowed", CellOf(Water, Sand, Grass), Up, CellOf(Water, Sand, Grass)},
		{"Cell Grass Left allowed", CellOf(Grass), Left, CellOf(Grass)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Allowed(tt.direction); got != tt.want {
				t.Errorf("Allowed(%v) on %04b = %04b, want %04b", tt.direction, tt.cell, got, tt.want)
			}
		})
	}
}
