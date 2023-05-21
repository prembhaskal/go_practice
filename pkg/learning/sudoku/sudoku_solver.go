package sudoku

import (
	"fmt"
)

// grid is divided into 9 parts, each part of 9 cells each
// [0, 1 , 2]
// [3, 4 , 5]
// [6, 7 , 8]
type Grid struct {
	ar [][]int
}

type Set map[int]int


// func SolveSingleCell(grid *Grid) {

// }


func FindElements(grid *Grid, partnum int) Set {
	
}


// returns the top,left cell for the given part
// panics for out of bounds part
func FindStartCell(partnum int) (int, int) {
	if partnum < 0 || partnum > 8 {
		panic(fmt.Sprintf("partnum out of bounds: %d", partnum))
	}


	row := partnum/3 // truncates towards 3
	row = row * 3    // gets the boundary

	col := (partnum % 3 ) * 3

	return row, col
}

// returns the part number this cell belongs too
// panic if the cell is out of bounds
func FindGridPart(row, col int) int {
	// boundary check
	if row < 0 || row > 8 {
		panic(fmt.Sprintf("out of bounds row: %d", row))
	}
	if col < 0 || col > 8 {
		panic(fmt.Sprintf("out of bounds col: %d", col))
	}

	colpart := col / 3
	rowpart := row / 3

	return colpart + 3*rowpart
}
