package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var primeRange int = 100000

var br *bufio.Reader
var bw *bufio.Writer

func main() {

	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)

	defer bw.Flush()

	primelist := initPrimes()

	fmt.Fprintf(bw, "total primes: %d\n", len(primelist))
	printList(primelist)

	var tc int
	fmt.Fscanf(br, "%d\n", &tc)
	fmt.Fprintf(bw, "testcases: %d\n", tc)

	for ;tc > 0; tc-- {
		var m, n int
		fmt.Fscanf(br, "%d %d\n", &m, &n)

		fmt.Fprintf(bw, "m: %d, n: %d\n", m, n)
	}

}

func initPrimes() []int {
	sqrt := int(math.Sqrt(float64(primeRange))) + 1

	// fmt.Printf("sqrt is %d", sqrt)

	primes := make([]bool, primeRange + 1)

	for idx, _ := range(primes) {
		primes[idx] = true
	}

	for i := 2; i <= sqrt; i++ {
		for j := i * i; j < len(primes) ; j = j + i {
			primes[j] = false
		}
	}

	pl := make([]int, 0)
	pl = append(pl, 2)

	for i := 3; i < len(primes); i = i + 2 {
		if primes[i] {
			pl = append(pl, i)
		}
	}

	return pl
}

func printList(primelist []int) {
	for _, val := range(primelist) {
		fmt.Fprintf(bw, "%d ", val)
	}
	fmt.Fprintln(bw)
}

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
// func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
// func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

// func main() {
//   // STDOUT MUST BE FLUSHED MANUALLY!!!
//   defer writer.Flush()

//   var a, b int
//   scanf("%d %d\n", &a, &b)
//   printf("%d\n", a+b)
// }
