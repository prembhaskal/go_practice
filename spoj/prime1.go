package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var primeRange int = 100000

func main() {

	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	defer bw.Flush()

	initPrimes()

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

	fmt.Printf("sqrt is %d", sqrt)

	primes := make([]bool, primeRange + 1)

	for idx, _ := range(primes) {
		primes[idx] = true
	}

	return []int{}
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
