package wfc

import (
	"math/rand/v2"
	"testing"
)

func TestCellHas(t *testing.T) {

	tests := []struct {
		name string
		cell Cell
		tile Tile
		want bool
	}{
		{"tile 0 present", 0b0001, Tile(0), true},
		{"tile 0 not present", 0b0000, Tile(0), false},
		{"tile 0 with others present", 0b1110, Tile(0), false},
		{"tile 1 present", 0b0011, Tile(1), true},
		{"tile 2 present", 0b0111, Tile(2), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Has(tt.tile); got != tt.want {
				t.Errorf("Has(%04b) on %04b = %v, want %v", tt.tile, tt.cell, got, tt.want)
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
		{"four ones", 0b1111, 4},
		{"empty", 0b0000, 0},
		{"2 Ones", 0b0011, 2},
		{"3 ones", 0b1101, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Count(); got != tt.want {
				t.Errorf("Count(%04b) = %v, want %v", tt.cell, got, tt.want)
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
		{"remove tile 0", 0b0001, Tile(0), 0b0000},
		{"remove tile 0 from nothing", 0b0000, Tile(0), 0b0000},
		{"remove tile 1", 0b1111, Tile(1), 0b1101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Without(tt.tile); got != tt.want {
				t.Errorf("Without(%v) on %04b = %04b, want %04b", tt.tile, tt.cell, got, tt.want)
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
		{"empty cell", []Tile{}, 0b000},
		{"tile 0 cell", []Tile{0}, 0b001},
		{"tile 0, 1", []Tile{0, 1}, 0b011},
		{"tile 1, 2", []Tile{1, 2}, 0b110},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CellOf(tt.of...); got != tt.want {
				t.Errorf("CellOf(%v) = %04b, want %04b", tt.of, got, tt.want)
			}
		})
	}
}

func TestCellCollapse(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	c := CellOf(0, 2)

	var seen Cell

	for range 100 {
		got := c.Collapse(r)
		if got.Count() != 1 {
			t.Fatalf("Collapse(%04b) = %04b, want exactly one tile", c, got)
		}
		if got&c != got {
			t.Fatalf("Collapse(%04b) = %04b, not a subset", c, got)
		}
		seen |= got
	}

	if seen != c {
		t.Errorf("after 100 collapses saw %04b, want all of %04b", seen, c)
	}
}

func TestCollapseEmptyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Collapse on empty celll did not panic")
		}
	}()

	r := rand.New(rand.NewPCG(1, 2))
	Cell(0).Collapse(r)
}
