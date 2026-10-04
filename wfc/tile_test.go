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

func TestTileSetAddReturnsSequentialIndices(t *testing.T) {
	var ts TileSet

	names := []string{"Water", "Sand", "Grass"}
	for want, name := range names {
		got := ts.Add(name, '?')
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
	tile := ts.Add("Water", '~')

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
		ts.Add("t", '?')
	}

	defer func() {
		if recover() == nil {
			t.Errorf("Add beyond 64 tiles did not panic")
		}
	}()
	ts.Add("one too many", '?')
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
	water := ts.Add("Water", '~')
	sand := ts.Add("Sand", '.')
	grass := ts.Add("Grass", '#')

	ts.Connect(water, water)
	ts.Connect(water, sand)
	ts.Connect(sand, grass)

	assertSymmetric(t, &ts)
}
