package sudoku_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/prembhaskal/go_practice/pkg/learning/sudoku"
)

var cells1 = [][]int{
	{0, 0, 0, 0, 1, 7, 3, 0, 0},
	{0, 0, 0, 0, 0, 8, 0, 2, 0},
	{8, 0, 6, 0, 0, 0, 0, 0, 5},

	{5, 0, 0, 0, 0, 6, 0, 4, 3},
	{0, 0, 2, 0, 0, 0, 7, 0, 0},
	{9, 3, 0, 2, 0, 0, 0, 0, 6},

	{7, 0, 0, 0, 0, 0, 5, 0, 8},
	{0, 2, 0, 5, 0, 0, 0, 0, 0},
	{0, 0, 8, 7, 3, 0, 0, 0, 0},
}

var cells2 = [][]int{
	{0, 0, 0, 0, 0, 9, 8, 0, 0},
	{0, 0, 0, 0, 0, 1, 0, 5, 7},
	{1, 0, 0, 0, 5, 0, 0, 0, 4},

	{0, 4, 6, 5, 0, 0, 0, 0, 0},
	{0, 2, 0, 0, 7, 0, 0, 3, 0},
	{0, 0, 0, 0, 0, 8, 9, 4, 0},

	{7, 0, 0, 0, 4, 0, 0, 0, 2},
	{4, 6, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 3, 9, 0, 0, 0, 0, 0},
}

var emptyGrid = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},

	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func TestFindGridParts(t *testing.T) {
	// TODO - check the panic part

	// normal checks

	testdata := []struct {
		r int
		c int
		p int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{1, 0, 0},
		{1, 1, 0},
		{1, 2, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 0},

		{0, 3, 1},
		{0, 4, 1},
		{0, 5, 1},
		{1, 3, 1},
		{1, 4, 1},
		{1, 5, 1},
		{2, 3, 1},
		{2, 4, 1},
		{2, 5, 1},

		{0, 6, 2},
		{0, 7, 2},
		{0, 8, 2},
		{1, 6, 2},
		{1, 7, 2},
		{1, 8, 2},
		{2, 6, 2},
		{2, 7, 2},
		{2, 8, 2},

		{3, 0, 3},
		{3, 1, 3},
		{3, 2, 3},
		{4, 0, 3},
		{4, 1, 3},
		{4, 2, 3},
		{5, 0, 3},
		{5, 1, 3},
		{5, 2, 3},

		{3, 3, 4},
		{3, 4, 4},
		{3, 5, 4},
		{4, 3, 4},
		{4, 4, 4},
		{4, 5, 4},
		{5, 3, 4},
		{5, 4, 4},
		{5, 5, 4},

		{3, 6, 5},
		{3, 7, 5},
		{3, 8, 5},
		{4, 6, 5},
		{4, 7, 5},
		{4, 8, 5},
		{5, 6, 5},
		{5, 7, 5},
		{5, 8, 5},

		{6, 6, 8},
		{6, 7, 8},
		{6, 8, 8},
		{7, 6, 8},
		{7, 7, 8},
		{7, 8, 8},
		{8, 6, 8},
		{8, 7, 8},
		{8, 8, 8},
	}

	for _, tt := range testdata {
		actpart := sudoku.FindGridPart(tt.r, tt.c)
		assert.Equal(t, tt.p, actpart)
	}
}

func TestFindStartCell(t *testing.T) {
	testdata := []struct {
		part int
		row  int
		col  int
	}{
		{0, 0, 0},
		{1, 0, 3},
		{2, 0, 6},
		{3, 3, 0},
		{4, 3, 3},
		{5, 3, 6},
		{6, 6, 0},
		{7, 6, 3},
		{8, 6, 6},
	}

	for _, tt := range testdata {
		ar, ac := sudoku.FindStartCell(tt.part)
		assert.Equal(t, ar, tt.row)
		assert.Equal(t, ac, tt.col)
	}
}

// TODO add proper asse
func TestFindElementsInPart(t *testing.T) {
	sampleGrid := &sudoku.Grid{
		Ar: cells1,
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("part: %d\n", i)
		elems := sudoku.FindElementsInPart(sampleGrid, i)
		fmt.Printf("elems are %v\n", elems)
	}

	for r := 0; r < 9; r++ {
		fmt.Printf("vals in row: %d\n", r)
		elems := sudoku.FindElementsInRow(sampleGrid, r)
		fmt.Printf("elems are %v\n", elems)
	}

	for c := 0; c < 9; c++ {
		fmt.Printf("vals in column: %d\n", c)
		elems := sudoku.FindElementsInCol(sampleGrid, c)
		fmt.Printf("elems are %v\n", elems)
	}
}

func TestMergeSets1(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells1}

	relems := sudoku.FindElementsInRow(sampleGrid, 0)
	celems := sudoku.FindElementsInCol(sampleGrid, 0)
	aelems := sudoku.FindElementsInPart(sampleGrid, 0)

	elems := sudoku.MergeSets1(relems, celems, aelems)

	expelems := map[int]int{
		0: 1, 1: 1, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1,
	}

	assert.Equal(t, elems, sudoku.Set(expelems))

	expmiss := map[int]int{2: 1}
	miss := sudoku.FindMissingInSet(elems)
	assert.Equal(t, sudoku.Set(expmiss), miss)
}

func TestSolveSingleCell(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells1}
	// sudoku.Debug = true
	sudoku.SolveSingleCell(sampleGrid)
}

func TestRemoveFromSet(t *testing.T) {
	allset := sudoku.GetAllElemsSet()
	sudoku.RemoveFromSet(allset, 1, 5)
	fmt.Printf("allset is %v\n", allset)
}

func TestUpdateCellsWithMeta(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells1}
	sudoku.UpdateCellsWithMeta(sampleGrid)

	printcell := func(row, col int) {
		cell := sampleGrid.Cells[row][col]
		if cell.Val != 0 {
			return
		}
		fmt.Println(cell)
	}

	executeForEachCell(sampleGrid, printcell)
}

func executeForEachCell(grid *sudoku.Grid, f func(int, int)) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			f(r, c)
		}
	}
}

func TestSolveSingleCellWithMeta(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells1}
	// sudoku.Debug = true
	sudoku.UpdateCellsWithMeta(sampleGrid)
	sudoku.SolveSingleCellUsingMeta(sampleGrid)
}

func TestSolveShadowFromAdjacentParts(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells1}
	// sudoku.Debug = true
	sudoku.UpdateCellsWithMeta(sampleGrid)
	sudoku.SolveAloneInPart(sampleGrid)
}

func TestSolveUsingCurrentTechs(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells2}
	// sudoku.Debug = true
	sudoku.UpdateCellsWithMeta(sampleGrid)
	sudoku.SolveSingleCellUsingMeta(sampleGrid)
	sudoku.SolveAloneInPart(sampleGrid)
	fmt.Println(sampleGrid)
}

func TestSolveWithUpdates(t *testing.T) {
	sampleGrid := &sudoku.Grid{Ar: cells2}

	i := 0
	for {
		i++
		// sudoku.Debug = true
		sudoku.UpdateCellsWithMeta(sampleGrid)
		// sudoku.AlonePairInPartUpdatesCheck(sampleGrid)
		// sudoku.UpdateValidsInGrid(sampleGrid)
		gridUpdate := sudoku.NewGridUpdate()

		sudoku.SingleCellUpdates(sampleGrid, gridUpdate)
		sudoku.AloneInPartUpdates(sampleGrid, gridUpdate)

		if len(gridUpdate) == 0 {
			fmt.Printf("no more moves found, total iterations: %d\n", i)
			// sudoku.AlonePairInPartUpdatesCheck(sampleGrid)
			fmt.Printf("final grid: \n%s\n", sampleGrid)
			break
		}

		sampleGrid.UpdateGrid(gridUpdate)

		fmt.Printf("current status: \n%s\n\n", sampleGrid)

		fmt.Printf("iteration: %d , solving ...\n", i)
		time.Sleep(250 * time.Millisecond)
	}
}
