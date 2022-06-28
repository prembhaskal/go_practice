package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("hello world\n")

	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("hello from routine\n")
			if i == 10 {
				panic("error in go routine")
			}
			i++
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("hello from main\n")
	}

}
