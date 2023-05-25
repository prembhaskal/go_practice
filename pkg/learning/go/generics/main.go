package main

import "fmt"

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non generic sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}


func SumIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
} 