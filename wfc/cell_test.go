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
		{"Remove sand", 0b1111, Sand, 0b1111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cell.Without(tt.tile); got != tt.want {
				t.Errorf("Has(%v) = %v, want %v", tt.tile, got, tt.want)
			}
		})
	}
}
