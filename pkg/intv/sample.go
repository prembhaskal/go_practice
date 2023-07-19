package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Printf("hello %s\n", "world")

	chars := []rune{'A', 'B', 'C', 'D'}

	var sb strings.Builder
	for i := 0; i < 1000; i++ {
		idx := rand.Intn(4)
		sb.WriteRune(chars[idx])
	}

	fmt.Println(sb.String())
}
