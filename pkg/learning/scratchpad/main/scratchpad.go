package main

import (
	"fmt"
	"os"
)

func main() {
	in := os.Stdin
	for i := 0; i < 10; i++ {
		var inp string
		fmt.Printf("your input: ")
		fmt.Fscan(in, &inp)
		fmt.Printf("you entered: %s\n", inp)
	}
}
