package sudoku

import (
	"fmt"
	"log"
)

// grid is divided into 9 parts, each part of 9 cells each
// [0, 1, 2]
// [3, 4, 5]
// [6, 7, 8]
type Grid struct {
	Ar    [][]int
	Cells [][]*Cell
}

type Cell struct {
	Row     int
	Col     int
	Valid   Set
	InValid Set
	Val     int
}

func (c Cell) String() string {
	return fmt.Sprintf("cell at row: %d, col: %d, Val: %d, Valid: %v", c.Row, c.Col, c.Val, c.Valid)
}

type Set map[int]int

var Debug = false

func UpdateCellsWithMeta(grid *Grid) {
	// invalid not used for now
	initCellsInGrid(grid)
	addEmptyCellInGrid(grid)
	updateValidsInGrid(grid)
}

func updateValidsInGrid(grid *Grid) {
	// for every non empty cell
	//   update valid cells in the partnum, row and col
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := grid.Cells[r][c]
			if cell.Val != 0 {
				updateValidInRow(grid, r, cell.Val)
				updateValidInCol(grid, c, cell.Val)
				updateValidInPart(grid, FindGridPart(r, c), cell.Val)
			}
		}
	}
}

func updateValidInRow(grid *Grid, row, remval int) {
	for c := 0; c < 9; c++ {
		cell := grid.Cells[row][c]
		if cell.Val == 0 {
			RemoveFromSet(cell.Valid, remval)
		}
	}
}

func updateValidInCol(grid *Grid, col, remval int) {
	for r := 0; r < 9; r++ {
		cell := grid.Cells[r][col]
		if cell.Val == 0 {
			RemoveFromSet(cell.Valid, remval)
		}
	}
}

func updateValidInPart(grid *Grid, part, remval int) {
	f := func(row, col int) {
		cell := grid.Cells[row][col]
		if cell.Val == 0 {
			RemoveFromSet(cell.Valid, remval)
		}
	}
	ExecuteFuncForCellsInPart(grid, part, f)
}

func ExecuteFuncForCellsInPart(grid *Grid, partnum int, f func(int, int)) {
	roff, coff := FindStartCell(partnum)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			f(r+roff, c+coff)
		}
	}
}

func RemoveFromSet(set Set, vals ...int) {
	for _, v := range vals {
		delete(set, v)
	}
}

func addEmptyCellInGrid(grid *Grid) {
	// for the Grid, initiate the Cell
	// for non empty cell, fill only val
	// for every empty cell, first fill the default valid
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := &Cell{
				Row: r,
				Col: c,
				Val: grid.Ar[r][c],
			}

			if cell.Val == 0 {
				cell.Valid = GetAllElemsSet()
			}

			grid.Cells[r][c] = cell
		}
	}
}

func initCellsInGrid(grid *Grid) {
	cells := make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		cells[i] = make([]*Cell, 9)
	}
	grid.Cells = cells
}

func GetAllElemsSet() Set {
	mp := make(map[int]int)
	for n := 1; n <= 9; n++ {
		mp[n] = 1
	}
	return mp
}

func SolveSingleCellUsingMeta(grid *Grid) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := grid.Cells[r][c]
			// check in each cell, if the map length is just 1
			if len(cell.Valid) == 1 {
				missval := FindFirstKey(cell.Valid)
				part := FindGridPart(r, c)
				fmt.Printf("lone miss at part:%d row:%d, col:%d val: %d\n", part, r, c, missval)
			}
		}
	}
}

// assuming meta for grid is already updated.
func SolveShadowFromAdjacentParts(grid *Grid) {
	for part := 0; part < 9; part++ {
		for num := 0; num < 9; num++ {

			validcnt := 0
			var lastValidCell *Cell
			roff, coff := FindStartCell(part)

			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					row := r + roff
					col := c + coff
					cell := grid.Cells[row][col]
					if cell.Valid[num] > 0 {
						validcnt++
						lastValidCell = cell
					}
				}
			}

			if validcnt == 1 {
				fmt.Printf("part: %d, num: %d valid only in cell: %s\n", part, num, lastValidCell)
			}
		}
	}
}

func executeForEachCell(grid Grid, f func(int, int)) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			f(r, c)
		}
	}
}

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

func FindMissingInSet(set Set) Set {
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

	row := partnum / 3 // truncates towards 3
	row = row * 3      // gets the boundary

	col := (partnum % 3) * 3

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
