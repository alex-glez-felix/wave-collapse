package wfc

import (
	"errors"
	"maps"
	"math/rand/v2"
	"testing"
)

func TestNewWave(t *testing.T) {
	tests := []struct {
		name          string
		width, height int
	}{
		{"1x1", 1, 1},
		{"wide 5x2", 5, 2},
		{"tall 2x5", 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWave(tt.width, tt.height)

			if w.width != tt.width || w.height != tt.height {
				t.Fatalf("size = %dx%d, want %dx%d", w.width, w.height, tt.width, tt.height)
			}
			if got, want := len(w.grid), tt.width*tt.height; got != want {
				t.Fatalf("len(grid) = %d, want %d", got, want)
			}
			for i, c := range w.grid {
				if c != AllTiles {
					t.Errorf("grid[%d] = %04b, want %04b", i, c, AllTiles)
				}
			}
		})
	}
}

func TestWave_lowestEntropy(t *testing.T) {
	three := CellOf(Water, Sand, Grass)
	two := CellOf(Water, Sand)
	one := CellOf(Water)

	tests := []struct {
		name   string
		grid   []Cell
		want   map[int]bool
		wantOK bool
	}{
		{"ties after higher", []Cell{three, two, two}, map[int]bool{1: true, 2: true}, true},
		{"single candidate", []Cell{one, three, one}, map[int]bool{1: true}, true},
		{"all collapsed", []Cell{one, one, one}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wave{grid: tt.grid, width: len(tt.grid), height: 1}
			r := rand.New(rand.NewPCG(1, 2))
			seen := map[int]bool{}

			for range 100 {
				i, ok := w.lowestEntropy(r)
				if ok != tt.wantOK {
					t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
				}
				if !ok {
					break
				}
				if !tt.want[i] {
					t.Fatalf("got index %d, want one of %v", i, tt.want)
				}
				seen[i] = true
			}

			if tt.wantOK && !maps.Equal(seen, tt.want) {
				t.Errorf("seen %v, want all of %v", seen, tt.want)
			}
		})
	}
}

func newRow(width int) *Wave {
	grid := make([]Cell, width)
	for i := range grid {
		grid[i] = AllTiles
	}
	return &Wave{grid: grid, width: width, height: 1}
}

func TestPropagateRestrictsNeighbors(t *testing.T) {
	for tile := range NumTiles {
		t.Run(tile.String(), func(t *testing.T) {
			w := newRow(3)
			w.grid[1] = CellOf(tile) // colapsamos el centro a mano

			if err := w.propagate(1); err != nil {
				t.Fatalf("propagate: unexpected error: %v", err)
			}

			if got, want := w.grid[0], AllTiles&rules[tile][Left]; got != want {
				t.Errorf("left = %04b, want %04b", got, want)
			}
			if got, want := w.grid[2], AllTiles&rules[tile][Right]; got != want {
				t.Errorf("right = %04b, want %04b", got, want)
			}
			if got, want := w.grid[1], CellOf(tile); got != want {
				t.Errorf("center changed: %04b, want %04b", got, want)
			}
		})
	}
}

func TestPropagateContradiction(t *testing.T) {
	// buscamos una ficha que NO pueda estar a la izquierda de Water
	for tile := range NumTiles {
		if rules[Water][Left].Has(tile) {
			continue
		}

		w := newRow(2)
		w.grid[0] = CellOf(tile)
		w.grid[1] = CellOf(Water)

		err := w.propagate(1)
		if !errors.Is(err, ErrContradiction) {
			t.Errorf("tile %v left of Water: err = %v, want ErrContradiction", tile, err)
		}
		return
	}
	t.Skip("every tile is allowed left of Water; no contradiction to test")
}
