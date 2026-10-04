package main

import (
	"fmt"
	"github.com/alex-glez-felix/wave-collapse/wfc"
)

type Position struct {
	X int
	Y int
}

func main() {
	fmt.Println("Hola mundo")

	fmt.Printf("Num Tiles: %v\n", wfc.NumTiles)
	fmt.Printf("Water: %v\n", wfc.Water)
	fmt.Printf("Sand: %v\n", wfc.Sand)
	fmt.Printf("Grass: %v\n", wfc.Grass)
}
