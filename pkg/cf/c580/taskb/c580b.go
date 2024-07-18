package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := newInputReader()
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(in, writer)
}

func solve(in *inputreader, out io.Writer) {
	n := in.nextInt()
	d := in.nextInt()

	frnds := make([][]int, n)
	for i := 0; i < n; i++ {
		frnds[i] = make([]int, 2)
		frnds[i][0] = in.nextInt()
		frnds[i][1] = in.nextInt()
	}

	sort.Slice(frnds, func(i, j int) bool {
		if frnds[i][0] < frnds[j][0] {
			return true
		}
		return false
	})

	bptr := 0
	fptr := 1
	total := frnds[0][1]
	maxtotal := total
	for ; fptr < n; fptr++ {
		for frnds[fptr][0]-frnds[bptr][0] >= d {
			total = total - frnds[bptr][1]
			bptr++
		}
		total = total + frnds[fptr][1]
		maxtotal = max(maxtotal, total)
	}

	fmt.Fprintln(out, maxtotal)
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
