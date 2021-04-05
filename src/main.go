package main

import (
	"fmt"

	"github.com/Mustafa-Dara-Ozevin/go-sudoku/src/game"
)

func main() {
	grid := game.NewGrid()
	grid.GenerateGrid()
	grid.Print()
	fmt.Println("Solving..")
	grid.Solve()
	grid.Print()
}
