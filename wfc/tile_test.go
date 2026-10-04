package wfc

import "testing"

func TestTileAllowed(t *testing.T) {
	tests := []struct {
		name      string
		tile      Tile
		direction Direction
		want      Cell
	}{
		{"Water up allowed", Water, Up, CellOf(Water, Sand)},
		{"Grass down allowed", Grass, Down, CellOf(Grass, Sand)},
		{"Sand down allowed", Sand, Down, CellOf(Grass, Sand, Water)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tile.Allowed(tt.direction); got != tt.want {
				t.Errorf("Allowed(%v) on %04b = %v, want %04b", tt.direction, tt.tile, got, tt.want)
			}
		})
	}
}
