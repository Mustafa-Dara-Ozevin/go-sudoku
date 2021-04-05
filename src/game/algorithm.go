package game

import (
	"math/rand"
	"time"
)

func (g Grid) isFull() bool {
	for _, tile := range g.Board {
		if tile == empty {
			return false
		}
	}
	return true
}

// fills the gird by placing random legal number to first empty square and recursively calling itself
// backtacks if there is no legal number for square
func (g *Grid) fill() bool {
	emptySqr := g.findEmpty()
	if emptySqr == -1 {
		return true
	}
	numbers := [...]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
	for _, num := range numbers {
		if g.setTile(emptySqr, num) {
			if g.fill() {
				return true
			}
			g.takeback(emptySqr)
		}
	}
	return false
}

// checks if grid has more than one solution by incrementing class variable count each time it fills
// the grid entirely doesn't actually solve the sudoku
func (g *Grid) checkGrid() bool {
	emptyTile := g.findEmpty()
	if emptyTile == -1 {
		return true
	}
	for num := uint8(1); num < 10; num++ {
		if g.setTile(emptyTile, num) {
			if g.isFull() {
				g.counter++
				break
			} else if g.checkGrid() {
				return true
			}
			g.takeback(emptyTile)
		}
	}
	return false
}

// Solves the sudoku by backtracking method for more information https://www.wikiwand.com/en/Sudoku_solving_algorithms
func (g *Grid) Solve() bool {
	emptyTile := g.findEmpty()
	if emptyTile == -1 {
		return true
	}

	for num := uint8(1); num < 10; num++ {
		if g.setTile(emptyTile, num) {

			if g.Solve() {
				return true
			}
			g.takeback(emptyTile)
		}
	}
	return false
}

// generates legal sudoku position by removing random number from grid and checking if its legal with checkGrid
// attempts parameter can be changed to adjust generation time and diffuculty of sudoku
func (g *Grid) GenerateGrid() bool {
	g.fill()
	attempts := 500
	g.counter = 1
	for attempts > 1 {
		row := rand.Intn(9)
		col := rand.Intn(9)
		cord := (row * 9) + col
		if g.Board[cord] != empty {
			deletedNum := g.Board[cord]
			g.takeback(cord)
			boardBackup := g.Board
			g.counter = 0
			g.checkGrid()
			g.Board = boardBackup
			if g.counter != 1 {
				g.setTile(cord, deletedNum)
			}

		}
		attempts--
	}
	return true
}
