package main

import (
	"fmt"
	"sort"
	// "time"
)

type emp struct {
	name string
	age  int
	id   string
}

func sortEmp() {
	emps := make([]emp, 0)

	sort.Slice(emps, func(i, j int) bool {
		return emps[i].age < emps[j].age
	})
}

func main() {
	fmt.Printf("hello world\n")

	num := 123

	// sumin := make(chan int)
	sumout := make(chan int)
	// go sum(sumin, sumout)

	sqin := make(chan int)
	sqsum := make(chan int)
	go squaresum(sqin, sqsum)

	cin := make(chan int)
	csum := make(chan int)
	go cubesum(cin, csum)

	go sum2(sqsum, csum, sumout)

	go func() {
		for num > 0 {
			digit := num % 10
			num = num / 10
			sqin <- digit
			cin <- digit
		}
		close(sqin)
		close(cin)
	}()

	totalans := <-sumout
	fmt.Printf("total is %d\n", totalans)

	// time.Sleep(5 * time.Second)
}

func sum(in, out chan int) {
	total := 0
	for sq := range in {
		fmt.Printf("sum in: %d\n", sq)
		total = total + sq
	}
	out <- total
}

func sum2(in1, in2, out chan int) {
	total := 0
out:
	for {
		select {
		case sq := <-in1:
			total = total + sq
		case cb := <-in2:
			total = total + cb
		default:
			break out
		}
	}

	out <- total
}

func squaresum(in, out chan int) {
	for num := range in {
		fmt.Printf("squaresum: %d\n", num)
		out <- num * num
	}
	close(out)
}

func cubesum(in, out chan int) {
	for num := range in {
		fmt.Printf("cubesum: %d\n", num)
		out <- num * num * num
	}
	close(out)
}

// num  123
// concurrent
// sum of squares of each num -
// sum of cubes of each num
// sum of squares + cubes
