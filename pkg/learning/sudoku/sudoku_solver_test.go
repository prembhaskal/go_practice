package sudoku_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/prembhaskal/go_practice/pkg/learning/sudoku"
)

func TestFindGridParts(t *testing.T) {
	// TODO - check the panic part

	// normal checks

	testdata := []struct {
		r int
		c int
		p int
	}{
		{0, 0, 0}, {0, 1, 0}, {0, 2, 0}, 
		{1, 0, 0}, {1, 1, 0}, {1, 2, 0}, 
		{2, 0, 0}, {2, 1, 0}, {2, 2, 0}, 
		
		{0, 3, 1}, {0, 4, 1}, {0, 5, 1}, 
		{1, 3, 1}, {1, 4, 1}, {1, 5, 1}, 
		{2, 3, 1}, {2, 4, 1}, {2, 5, 1}, 
	
		{0, 6, 2}, {0, 7, 2}, {0, 8, 2}, 
		{1, 6, 2}, {1, 7, 2}, {1, 8, 2}, 
		{2, 6, 2}, {2, 7, 2}, {2, 8, 2}, 
		

		{3, 0, 3}, {3, 1, 3}, {3, 2, 3}, 
		{4, 0, 3}, {4, 1, 3}, {4, 2, 3}, 
		{5, 0, 3}, {5, 1, 3}, {5, 2, 3}, 
		
		{3, 3, 4}, {3, 4, 4}, {3, 5, 4}, 
		{4, 3, 4}, {4, 4, 4}, {4, 5, 4}, 
		{5, 3, 4}, {5, 4, 4}, {5, 5, 4}, 
		
		{3, 6, 5}, {3, 7, 5}, {3, 8, 5}, 
		{4, 6, 5}, {4, 7, 5}, {4, 8, 5}, 
		{5, 6, 5}, {5, 7, 5}, {5, 8, 5},

		{6, 6, 8}, {6, 7, 8}, {6, 8, 8}, 
		{7, 6, 8}, {7, 7, 8}, {7, 8, 8}, 
		{8, 6, 8}, {8, 7, 8}, {8, 8, 8}, 
	}

	for _, tt := range testdata {
		actpart := sudoku.FindGridPart(tt.r, tt.c)
		assert.Equal(t, tt.p, actpart)
	}
}

func TestFindStartCell(t *testing.T) {
	testdata := []struct{
		part int
		row int
		col int
	}{
		{0, 0, 0}, {1, 0, 3}, {2, 0, 6},
		{3, 3, 0}, {4, 3, 3}, {5, 3, 6},
		{6, 6, 0}, {7, 6, 3}, {8, 6, 6},
	}

	for _, tt := range testdata {
		ar, ac := sudoku.FindStartCell(tt.part)
		assert.Equal(t, ar, tt.row)
		assert.Equal(t, ac, tt.col)
	}
}
