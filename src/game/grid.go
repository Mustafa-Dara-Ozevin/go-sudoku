package game

import (
	"fmt"
)

const empty = 0

// Initializer func
func NewGrid() Grid {
	var grid Grid
	grid.counter = 0
	for i := range grid.Board {
		grid.Board[i] = empty
	}
	return grid
}

// Util funcs
func rowFromCord(cord int) int {
	return cord / 9
}

func colFromCord(cord int) int {
	return cord % 9
}

type Grid struct {
	Board   [81]uint8
	counter int // used for counting possible solutions in checkGrid()
}

func (g *Grid) FromArray(gameArray [81]uint8) {
	g.Board = gameArray
}

// prints the board while highlighting each square
// Bug prints an extra "║" to the newline at the end
func (g Grid) Print() {
	fmt.Print("=====================================\n║ ")
	for i, tile := range g.Board {
		if tile == empty {
			fmt.Print("  ")
		} else {
			fmt.Printf("%d ", tile)

		}
		if (i+1)%3 != 0 {
			fmt.Print("| ")
		}
		if (i+1)%9 == 0 {
			if ((i+1)/9)%3 == 0 {
				fmt.Print("║\n=====================================\n║ ")
			} else {
				fmt.Print("║\n-------------------------------------\n║ ")
			}

		} else if (i+1)%3 == 0 {
			fmt.Print("║ ")
		}
	}
}

// takes where it's going to placed and number returns if it's legal or not
func (g *Grid) isLegal(cord int, num uint8) bool {
	// gets row from cord
	row, col := rowFromCord(cord), colFromCord(cord)
	// check if tile empty
	if g.Board[cord] != empty {
		return false
	}
	// places the number for checks needs to be removed before returning
	g.Board[cord] = num
	//check for column
	for rowIdx := 0; rowIdx >= 0 && rowIdx < 81; rowIdx += 9 {
		if row != rowIdx/9 && g.Board[rowIdx+col] == g.Board[cord] {
			g.Board[cord] = empty
			return false
		}
	}
	//check for row
	for colIdx := 0; colIdx < 9; colIdx++ {
		if col != colIdx && g.Board[row*9+colIdx] == g.Board[cord] {
			g.Board[cord] = empty
			return false
		}
	}
	// check for square
	// dividing by 3 and casting to int gives us what square the row/col is on starting by zero
	// for example 8/3 2.666... casting it to int gives us 2
	sqRow := int(row / 3)
	sqCol := int(col / 3)
	// multiplying sqRow by 27 gives us start row because square row changes each 27 tiles
	for rowIdx := sqRow * 27; rowIdx >= 0 && rowIdx < sqRow*27+27; rowIdx += 9 {
		// multiplying sqCol by 3 gives us start col because square row changes each 3 tiles
		for colIdx := sqCol * 3; colIdx < sqCol*3+3; colIdx++ {
			// dividing rowIdx by 9 when comparing to row because row is what row it is 0,1...8 etc
			// while rowIdx is actual cord of square 0,9...71 etc
			if !(rowIdx/9 == row && colIdx == col) && g.Board[rowIdx+colIdx] == g.Board[cord] {
				g.Board[cord] = empty
				return false
			}
		}
	}

	g.Board[cord] = empty
	return true
}

func (g *Grid) setTile(cord int, num uint8) bool {
	if g.isLegal(cord, num) {
		g.Board[cord] = uint8(num)
		return true
	}
	return false
}
func (g *Grid) takeback(cord int) bool {
	if g.Board[cord] != empty {
		g.Board[cord] = empty
		return true
	}
	return false
}
func (g Grid) findEmpty() int {
	for i := range g.Board {
		if g.Board[i] == empty {
			return i
		}
	}
	return -1
}
