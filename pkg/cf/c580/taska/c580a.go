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

// https://codeforces.com/contest/580/problem/A
func solve(in *inputreader, out io.Writer) {
	n := in.nextInt()
	ar := make([]int, n)
	for i := 0; i < n; i++ {
		ar[i] = in.nextInt()
	}

	maxlen := 1
	clen := 1
	prev := ar[0]
	for i := 1; i < n; i++ {
		if ar[i] >= prev {
			clen++
			maxlen = max(maxlen, clen)
		} else {
			clen = 1
		}
		prev = ar[i]
	}

	fmt.Fprintln(out, maxlen)
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
