package sudoku

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func NewSet(ar []int) Set {
	set := make(map[int]int)
	for _, v := range ar {
		set[v] = 1
	}
	return set
}

var Debug = false

type CellPos struct {
	Row int
	Col int
}

type GridUpdate map[CellPos]int

func NewGridUpdate() GridUpdate {
	return make(map[CellPos]int)
}

func (g GridUpdate) AddEntry(row, col, val int) {
	cellPos := CellPos{row, col}
	g[cellPos] = val
}

func (g *Grid) String() string {
	var sb strings.Builder
	sb.WriteString(g.Printline())
	for r := 0; r < 9; r++ {
		if r == 3 || r == 6 {
			sb.WriteString(g.Printline())
		}
		sb.WriteString("| ")
		for c := 0; c < 9; c++ {
			if c == 3 || c == 6 {
				sb.WriteString("| ")
			}
			sb.WriteString(strconv.Itoa(g.Ar[r][c]))
			sb.WriteString(" |")
		}
		sb.WriteString("\n")
	}
	sb.WriteString(g.Printline())
	return sb.String()
}

func (g *Grid) Printline() string {
	var sb strings.Builder
	sb.WriteString(" ")
	for i := 0; i < 18; i++ {
		sb.WriteString(" _")
	}
	sb.WriteString("\n")
	return sb.String()
}

func (g *Grid) UpdateGrid(gridUpdate GridUpdate) {
	for k, v := range gridUpdate {
		row := k.Row
		col := k.Col
		g.Ar[row][col] = v
	}
}

func UpdateCellsWithMeta(grid *Grid) {
	// invalid not used for now
	initCellsInGrid(grid)
	addEmptyCellInGrid(grid)
	updateValidsInGrid(grid)
	AlonePairInPartUpdatesCheck(grid)
}

// func UpdateValidsInGrid(grid *Grid) {
// 	updateValidsInGrid(grid)
// }

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

func updateValidInRowExcl(grid *Grid, row, remval int, excludeCols Set) {
	for c := 0; c < 9; c++ {
		cell := grid.Cells[row][c]
		if cell.Val == 0 && excludeCols[c] == 0 {
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

func updateValidInColExcl(grid *Grid, col, remval int, excludeRows Set) {
	for r := 0; r < 9; r++ {
		cell := grid.Cells[r][col]
		if cell.Val == 0 && excludeRows[r] == 0 {
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
	gridUpdate := make(map[CellPos]int)
	SingleCellUpdates(grid, gridUpdate)
}

func SingleCellUpdates(grid *Grid, gridUpdate GridUpdate) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := grid.Cells[r][c]
			// check in each cell, if the map length is just 1
			if len(cell.Valid) == 1 {
				missval := FindFirstKey(cell.Valid)
				part := FindGridPart(r, c)
				// fmt.Printf("LONE at part:%d row:%d, col:%d val: %d\n", part, r, c, missval)
				PrintMiss("LONE", part, r, c, missval)
				gridUpdate.AddEntry(r, c, missval)
			}
		}
	}
}

// assuming meta for grid is already updated.
func SolveAloneInPart(grid *Grid) {
	AloneInPartUpdates(grid, make(map[CellPos]int))
}

// check if a part, there exists a cell which has only one valid number.
func AloneInPartUpdates(grid *Grid, gridUpdate GridUpdate) {
	for part := 0; part < 9; part++ {
		for num := 1; num <= 9; num++ {

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
				// fmt.Printf("SHADOW part: %d, num: %d valid only in cell: %s\n", part, num, lastValidCell)
				// fmt.Printf(" at part:%d row:%d, col:%d val: %d\n", part, r, c, missval)
				PrintMiss("SHADOW", part, lastValidCell.Row, lastValidCell.Col, num)
				gridUpdate.AddEntry(lastValidCell.Row, lastValidCell.Col, num)
			}
		}
	}
}

// checks if two cells in same row or same column, have same number in the valid,
// and that number does not exist as part of valid in same cell.
func AlonePairInPartUpdatesCheck(grid *Grid) {
	for part := 0; part < 9; part++ {
		for num := 1; num <= 9; num++ {
			validcnt := 0
			var validCells []*Cell
			roff, coff := FindStartCell(part)

			for r := 0; r < 3; r++ {
				for c := 0; c < 3; c++ {
					row := r + roff
					col := c + coff
					cell := grid.Cells[row][col]
					if cell.Valid[num] > 0 {
						validcnt++
						validCells = append(validCells, cell)
					}
				}
			}

			if validcnt == 2 {
				// fmt.Printf("SHADOW part: %d, num: %d valid only in cell: %s\n", part, num, lastValidCell)
				// fmt.Printf(" at part:%d row:%d, col:%d val: %d\n", part, r, c, missval)
				// PrintMiss("SHADOW", part, lastValidCell.Row, lastValidCell.Col, num)
				// gridUpdate.AddEntry(lastValidCell.Row, lastValidCell.Col, num)
				cell1 := validCells[0]
				cell2 := validCells[1]
				if cell1.Col == cell2.Col || cell1.Row == cell2.Row {
					PrintMiss("ALONE PAIR PART 1", part, cell1.Row, cell1.Col, num)
					PrintMiss("ALONE PAIR PART 2", part, cell2.Row, cell2.Col, num)
					if cell1.Col == cell2.Col {
						// remove num from other rows of this column
						exclRow := NewSet([]int{cell1.Row, cell2.Row})
						updateValidInColExcl(grid, cell1.Col, num, exclRow)
					} else {
						// remove num from other cols of this row
						exclCol := NewSet([]int{cell1.Col, cell2.Col})
						updateValidInRowExcl(grid, cell1.Row, num, exclCol)
					}
				}
			}
		}
	}
}

func PrintMiss(misstype string, part, row, col, missval int) {
	fmt.Printf("%s at part:%d row:%d, col:%d val: %d\n", misstype, part, row, col, missval)
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
