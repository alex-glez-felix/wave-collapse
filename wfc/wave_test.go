package wfc

import (
	"maps"
	"math/rand/v2"
	"testing"
)

func TestNewWave(t *testing.T) {
	ts, _, _, _ := testTileSet()
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
			w := NewWave(tt.width, tt.height, ts)

			if w.width != tt.width || w.height != tt.height {
				t.Fatalf("size = %dx%d, want %dx%d", w.width, w.height, tt.width, tt.height)
			}
			if got, want := len(w.grid), tt.width*tt.height; got != want {
				t.Fatalf("len(grid) = %d, want %d", got, want)
			}
			for i, c := range w.grid {
				if c != ts.All() {
					t.Errorf("grid[%d] = %04b, want %04b", i, c, ts.All())
				}
			}
		})
	}
}

func TestWave_lowestEntropy(t *testing.T) {

	ts, water, sand, grass := testTileSet()

	three := CellOf(water, sand, grass)
	two := CellOf(water, sand)
	one := CellOf(water)

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
			w := &Wave{grid: tt.grid, width: len(tt.grid), height: 1, tiles: ts}
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
