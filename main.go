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

	ts := wfc.TileSet{}
	water := ts.Add("water", '~')
	sand := ts.Add("sand", '.')
	grass := ts.Add("grass", '#')
	ts.Connect(water, water)
	ts.Connect(water, sand)
	ts.Connect(sand, sand)
	ts.Connect(grass, sand)
	ts.Connect(grass, grass)

	for attempt := range 100 {
		wave := wfc.NewWave(20, 20, &ts)
		if err := wave.Run(r); err != nil {
			continue
		}
		fmt.Printf("Attempt: %v\n", attempt)
		fmt.Println(wave)
		return
	}
	fmt.Println("no solution after 100 attempts")

}
