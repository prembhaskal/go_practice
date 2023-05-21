package sudoku

import (
	"fmt"
	"log"
)

// grid is divided into 9 parts, each part of 9 cells each
// [0, 1 , 2]
// [3, 4 , 5]
// [6, 7 , 8]
type Grid struct {
	Ar [][]int
}

type Set map[int]int

var Debug = false

func SolveSingleCell(grid *Grid) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := grid.Ar[r][c]

			// check only empty cells
			if val != 0 {
				continue 
			}

			relems := FindElementsInRow(grid, r)
			celems := FindElementsInCol(grid, c)
			partnum := FindGridPart(r, c)
			pelems := FindElementsInPart(grid, partnum)

			aelems := MergeSets1(relems, celems, pelems)

			melems := FindMissingInSet(aelems)

			if Debug {
				log.Printf("row: %d, col: %d, miss: %v", r, c, melems)
			}

			if len(melems) == 1 {
				missval := FindFirstKey(melems)
				fmt.Printf("lone miss at part:%d row:%d, col:%d val: %d\n", partnum, r, c, missval)
			}
		}
	}
}

func FindFirstKey(set Set) int {
	for k, _ := range set {
		return k
	}
	panic("empty set")
}

func FindMissingInSet(set Set) Set{
	miss := make(map[int]int)
	for n := 1; n <= 9; n++ {
		if _, ok := set[n]; !ok {
			miss[n] = 1
		}
	}

	return miss
}

func MergeSets1(sets ...Set) Set {
	elems := make(map[int]int)
	for _, set := range sets {
		for k, _ := range set {
			elems[k] = 1
		}
	}
	return elems
}

func FindElementsInRow(grid *Grid, rownum int) Set {
	elems := make(map[int]int)
	for c := 0; c < 9; c++ {
		val := grid.Ar[rownum][c]
		elems[val]++
	}

	return elems
}

func FindElementsInCol(grid *Grid, col int) Set {
	elems := make(map[int]int)
	for r := 0; r < 9; r++ {
		val := grid.Ar[r][col]
		elems[val]++
	}

	return elems
}


func FindElementsInPart(grid *Grid, partnum int) Set {
	roff, coff := FindStartCell(partnum)
	elems := make(map[int]int)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			row := r + roff
			col := c + coff
			val := grid.Ar[row][col]
			elems[val]++	
		}
	}

	return elems
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
