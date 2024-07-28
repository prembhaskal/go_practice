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
	in := newInputReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solve(in, writer)
}

func solve(in *inputreader, out io.Writer) {
	n := in.nextInt()
	m := in.nextInt()
	k := in.nextInt()
	data := make(map[string][][]int)
	planets := make([]string, n)
	for i := 0; i < n; i++ {
		pname := in.next()
		planets[i] = pname
		pdata := make([][]int, m)
		for j := 0; j < m; j++ {
			buy := in.nextInt()
			sell := in.nextInt()
			count := in.nextInt()
			itemdata := []int{buy, sell, count}
			pdata[j] = itemdata
		}
		data[pname] = pdata
	}
	maxProfit := -1000000

	for i := 0; i < n; i++ {
		for j := 0 ; j < n; j++ {
			if i == j {
				continue
			}
			profitData := make([][]int, m)
			buydata := data[planets[i]]
			selldata := data[planets[j]]
			for item := 0; item < m; item++ {
				profitData[item] = []int{selldata[item][1] - buydata[item][0], buydata[item][1], buydata[item][2]}
			}
			// fmt.Printf("buy from: %s, sell to: %s, profit: %d\n", planets[i], planets[j], maxProfitBuySell(profitData, k))
			maxProfit = max(maxProfit, maxProfitBuySell(profitData, k))
		}
	}

	fmt.Fprintln(out, maxProfit)
}

// a[i][0] = profit of ith item
// a[i][1] = buy price of ith item
// a[i][2] = count of ith item
func maxProfitBuySell(a [][]int, limit int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] > a[j][0]
	})

	sell := 0
	// buy := 0
	for i := 0; i < len(a); i++ {
		if (a[i][0] > 0 ) {
			take := min(a[i][2], limit)
			sell = sell + take * a[i][0]
			// buy = buy + take * a[i][1]
			limit = limit - take
			if limit == 0 {
				break
			}
		}
	}

	return sell
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
