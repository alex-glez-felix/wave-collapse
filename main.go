package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/alex-glez-felix/wave-collapse/wfc"
)

type Position struct {
	X int
	Y int
}

func main() {
	s1 := rand.Uint64()
	s2 := rand.Uint64()
	fmt.Printf("s1: %v, s2: %v\n", s1, s2)
	r := rand.New(rand.NewPCG(s1, s2))

	defs := []struct {
		name  string
		glyph rune
		color uint8
	}{
		{"deep", '~', 18},
		{"water", '~', 33},
		{"sand", '.', 229},
		{"grass", '"', 70},
		{"forest", '♣', 22},
		{"mountain", '^', 244},
		{"snow", '*', 255},
	}

	var ts wfc.TileSet
	var prev wfc.Tile
	for i, d := range defs {
		tile := ts.Add(d.name, d.glyph, d.color)
		ts.Connect(tile, tile)
		if i > 0 {
			ts.Connect(prev, tile)
		}
		prev = tile
	}

	for attempt := range 100 {
		wave := wfc.NewWave(50, 50, &ts)
		if err := wave.Run(r); err != nil {
			continue
		}
		fmt.Printf("Attempt: %v\n", attempt)
		fmt.Println(wave.Render())
		return
	}
	fmt.Println("no solution after 100 attempts")

}
