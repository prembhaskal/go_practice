package main

// https://www.spoj.com/problems/PRIME1/
// 1 <= m <= n <= 1000000000, n-m <= 100000

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var primeRange int = 100000

var primeList []int

var br *bufio.Reader
var bw *bufio.Writer

func main() {

	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)

	defer bw.Flush()

	primeList = initPrimes()

	// fmt.Fprintf(bw, "total primes: %d\n", len(primelist))
	// printList()

	var tc int
	fmt.Fscanf(br, "%d\n", &tc)
	// fmt.Fprintf(bw, "testcases: %d\n", tc)

	for ; tc > 0; tc-- {
		var m, n int
		fmt.Fscanf(br, "%d %d\n", &m, &n)

		// fmt.Fprintf(bw, "m: %d, n: %d\n", m, n)

		rangeList := getRangePrimes(m, n)
		for _, v := range rangeList {
			fmt.Fprintf(bw, "%d\n", v)
		}
		fmt.Fprintln(bw)
	}

}

func initPrimes() []int {
	sqrt := int(math.Sqrt(float64(primeRange))) + 1

	// fmt.Printf("sqrt is %d", sqrt)

	primes := make([]bool, primeRange+1)

	for idx, _ := range primes {
		primes[idx] = true
	}

	for i := 2; i <= sqrt; i++ {
		for j := i * i; j < len(primes); j = j + i {
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

func printList() {
	for _, val := range primeList {
		fmt.Fprintf(bw, "%d ", val)
	}
	fmt.Fprintln(bw)
}

func getRangePrimes(start, end int) []int {
	/*
		start
		>= x1*prime1, (x1+1)*prime1, .... (mark as non prime)  where x1 > 1
		>= (x2)*prime2, (x2+1)*prime2, .... (mark as non prime) where x2 > 1

		end

		// all items marked as non prime, add in list

	*/

	numlen := end - start + 1

	boolList := make([]bool, numlen)

	for idx, _ := range boolList {
		boolList[idx] = true
	}

	if start == 1 {
		boolList[0] = false
	}

	for _, prime := range primeList {

		if prime*prime > end {
			break
		}

		// find smallest num such that num >= start and num % prime == 0
		startNum := start
		if start%prime != 0 {
			startNum = start + prime - (start % prime)
		}

		if prime == startNum {
			startNum = startNum + prime
		}

		for ; startNum <= end; startNum = startNum + prime {
			boolList[startNum-start] = false
		}

		// fmt.Printf("prime %d, primelist: %v\n", prime, boolList)
	}

	rangePrimeList := make([]int, 0)
	for idx, _ := range boolList {
		if boolList[idx] {
			rangePrimeList = append(rangePrimeList, idx+start)
		}
	}

	return rangePrimeList
}
