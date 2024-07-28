package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	fil, err := os.Open("C:\\Users\\pbhas\\test.txt")
	if err != nil {
		t.Errorf("%v", err)
		t.FailNow()
	}
	in := newInputReader(fil)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(in, writer)
}