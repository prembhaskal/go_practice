package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			<-ticker.C
			fmt.Printf("ticker invoked\n")
		}
	}()

	wg.Wait()
}
