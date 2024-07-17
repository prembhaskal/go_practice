package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	in := newInputReader()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(in, writer)
}
// https://codeforces.com/contest/1989/problem/C
func solve(in *inputreader, out io.Writer) {
	tests := in.nextInt()
	for i := 0; i < tests; i++ {
		n := in.nextInt()
		A := make([]int, n)
		for  j := 0; j < n; j++ {
			A[j] = in.nextInt()
		}
		B := make([]int, n)
		for j := 0; j < n; j++ {
			B[j] = in.nextInt()
		}
		fmt.Fprintln(out, findMaxRating(n, A, B))
	}
}

func findMaxRating(n int, A, B []int) int {
	totalA := 0
	totalB := 0
	for i := 0; i < n; i++ {
		if A[i] != B[i] {
			if A[i] > B[i] {
				totalA += A[i]
			} else {
				totalB += B[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if A[i] == B[i] {
			if A[i] == -1 {
				if totalA > totalB {
					totalA--
				} else {
					totalB--
				}
			} else if A[i] == 1 {
				if totalA > totalB {
					totalB++
				} else {
					totalA++
				}
			}
		}
	}
	return min(totalA, totalB)
}


type inputreader struct {
	reader *bufio.Scanner
}

func newInputReader() *inputreader {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	return &inputreader{
		reader: scn,
	}
}

func (i *inputreader) next() string {
	ok := i.reader.Scan()
	if !ok {
		panic("no data to read")
	}
	return i.reader.Text()
}

func (i *inputreader) nextraw() []byte {
	ok := i.reader.Scan()
	if !ok {
		panic("no data to read")
	}
	return i.reader.Bytes()
}

func (i *inputreader) nextInt() int {
	data := i.next()
	v, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		panic("error parsing int -> " + data)
	}
	return int(v)
}
func (i *inputreader) nextInt64() int64 {
	data := i.next()
	v, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		panic("error parsing int64 -> " + data)
	}
	return v
}

func (i *inputreader) nextFloat() float64 {
	data := i.next()
	v, err := strconv.ParseFloat(data, 64)
	if err != nil {
		panic("error parsing float -> " + data)
	}
	return v
}