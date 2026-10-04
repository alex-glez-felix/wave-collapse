package wfc

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
)

var ErrContradiction = errors.New("wfc: contradiction")

type Wave struct {
	grid   []Cell
	width  int
	height int
}

func NewWave(width, height int) *Wave {

	grid := make([]Cell, width*height)
	for i := range width * height {
		grid[i] = AllTiles
	}

	return &Wave{
		grid:   grid,
		width:  width,
		height: height,
	}
}

func (w *Wave) lowestEntropy(r *rand.Rand) (int, bool) {

	minEntropy := math.MaxInt
	var indexes []int

	for i, cell := range w.grid {
		if entropy := cell.Count(); entropy > 1 && entropy <= minEntropy {
			if entropy < minEntropy {
				indexes = indexes[:0]
				minEntropy = entropy
			}
			indexes = append(indexes, i)
		}
	}

	if len(indexes) == 0 {
		return 0, false
	}

	lowest := indexes[r.IntN(len(indexes))]
	return lowest, true
}

func (w *Wave) collapse(r *rand.Rand, index int) {
	w.grid[index] = w.grid[index].Collapse(r)
}

func (w *Wave) propagate(index int) error {
	stack := []int{index}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cell := w.grid[current]
		for direction := range NumDirections {
			n, ok := w.neighbor(current, direction)
			if !ok {
				continue
			}

			before := w.grid[n]
			intersected := before & cell.Allowed(direction)

			if intersected == 0 {
				return fmt.Errorf("cell %d emptied by cell %d: %w", n, current, ErrContradiction)
			}

			if before != intersected {
				w.grid[n] = intersected
				stack = append(stack, n)
			}
		}
	}
	return nil
}

var offsets = [NumDirections]struct{ dx, dy int }{
	Up:    {0, -1},
	Down:  {0, 1},
	Left:  {-1, 0},
	Right: {1, 0},
}

func (w *Wave) index(x, y int) int {
	return y*w.width + x
}

func (w *Wave) coords(index int) (x, y int) {
	x = index % w.width
	y = index / w.width
	return x, y
}

func (w *Wave) inBounds(x, y int) bool {
	return x >= 0 && x < w.width && y >= 0 && y < w.height
}

func (w *Wave) neighbor(i int, d Direction) (int, bool) {
	x, y := w.coords(i)
	offset := offsets[d]
	nx, ny := x+offset.dx, y+offset.dy

	if !w.inBounds(nx, ny) {
		return 0, false
	}
	return w.index(nx, ny), true
}
