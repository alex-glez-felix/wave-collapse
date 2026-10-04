package wfc

import (
	"fmt"
	"testing"
)

func testTileSet() (*TileSet, Tile, Tile, Tile) {
	var ts TileSet
	water := ts.Add("water", '~', 0)
	sand := ts.Add("sand", '.', 0)
	grass := ts.Add("grass", '#', 0)
	ts.Connect(water, water)
	ts.Connect(water, sand)
	ts.Connect(sand, sand)
	ts.Connect(sand, grass)
	ts.Connect(grass, grass)
	return &ts, water, sand, grass
}

func TestTileAllowed(t *testing.T) {
	ts, water, sand, grass := testTileSet()

	tests := []struct {
		name      string
		tile      Tile
		direction Direction
		want      Cell
	}{
		{"Water up allowed", water, Up, CellOf(water, sand)},
		{"Grass down allowed", grass, Down, CellOf(grass, sand)},
		{"Sand down allowed", sand, Down, CellOf(grass, sand, water)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ts.Allowed(tt.tile, tt.direction); got != tt.want {
				t.Errorf("Allowed(%v) on %04b = %v, want %04b", tt.direction, tt.tile, got, tt.want)
			}
		})
	}
}

func TestTileSetAddReturnsSequentialIndices(t *testing.T) {
	var ts TileSet

	names := []string{"Water", "Sand", "Grass"}
	for want, name := range names {
		got := ts.Add(name, '?', 0)
		if got != Tile(want) {
			t.Errorf("Add(%q) = %d, want %d", name, got, want)
		}
	}

	if got := len(ts.tiles); got != len(names) {
		t.Errorf("len(tiles) = %d, want %d", got, len(names))
	}
}

func TestTileSetAddStoresInfo(t *testing.T) {
	var ts TileSet
	tile := ts.Add("Water", '~', 0)

	info := ts.tiles[tile]
	if info.name != "Water" || info.glyph != '~' {
		t.Errorf("tiles[%d] = {%q, %q}, want {\"Water\", '~'}", tile, info.name, info.glyph)
	}
	for d := range NumDirections {
		if info.rules[d] != 0 {
			t.Errorf("rules[%v] = %04b, want 0", d, info.rules[d])
		}
	}
}

func TestTileSetAddPanicsOverLimit(t *testing.T) {
	var ts TileSet
	for range 64 {
		ts.Add("t", '?', 0)
	}

	defer func() {
		if recover() == nil {
			t.Errorf("Add beyond 64 tiles did not panic")
		}
	}()
	ts.Add("one too many", '?', 0)
}

func assertSymmetric(t *testing.T, ts *TileSet) {
	t.Helper()
	for a := range ts.tiles {
		for b := range ts.tiles {
			for d := range NumDirections {
				ab := ts.tiles[a].rules[d].Has(Tile(b))
				ba := ts.tiles[b].rules[d.Opposite()].Has(Tile(a))
				if ab != ba {
					t.Errorf("asymmetric: %d allows %d at %v = %v, but %d allows %d at %v = %v",
						a, b, d, ab, b, a, d.Opposite(), ba)
				}
			}
		}
	}
}

func TestTileSetConnectSymmetry(t *testing.T) {
	var ts TileSet
	water := ts.Add("Water", '~', 0)
	sand := ts.Add("Sand", '.', 0)
	grass := ts.Add("Grass", '#', 0)

	ts.Connect(water, water)
	ts.Connect(water, sand)
	ts.Connect(sand, grass)

	assertSymmetric(t, &ts)
}

func TestTileSetAll(t *testing.T) {
	tests := []struct {
		n    int
		want Cell
	}{
		{0, 0},
		{1, 0b1},
		{4, 0b1111},
		{64, ^Cell(0)},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			var ts TileSet
			for range tt.n {
				ts.Add("t", '?', 0)
			}
			if got := ts.All(); got != tt.want {
				t.Errorf("All() with %d tiles = %b, want %b", tt.n, got, tt.want)
			}
		})
	}
}
