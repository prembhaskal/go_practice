package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	in := newInputReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(in, writer)
}

func solve(in *inputreader, out io.Writer) {
	n := in.nextInt()
	maxCats = in.nextInt()
	cats = make([]int, n+1)
	for i := 1; i <= n; i++ {
		cats[i] = in.nextInt()
	}

	tree = make([][]int, n+1)
	for e := 1; e < n; e++ { // n-1 edges
		from := in.nextInt()
		to := in.nextInt()

		tree[from] = append(tree[from], to)
		tree[to] = append(tree[to], from)
	}
	visited = make([]bool, n+1)

	visitR(1, 0)

	fmt.Fprintln(out, ans)
}

var tree [][]int
var cats []int
var visited []bool
var ans int
var maxCats int

// ret rest count start at node
func visitR(node, prevCats int) {
	// fmt.Printf("visiting node: %d\n", node)
	visited[node] = true

	prevCats++
	if cats[node] == 0 {
		prevCats = 0
	}

	if prevCats > maxCats {
		return
	}

	if node != 1 && len(tree[node]) == 1 { // leaf
		ans++
		return
	}

	for _, next := range tree[node] {
		if !visited[next] {
			visitR(next, prevCats)
		}
	}

}

type inputreader struct {
	reader *bufio.Scanner
}

func newInputReader(r io.Reader) *inputreader {
	scn := bufio.NewScanner(r)
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
