// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")

	done := make(chan bool)

	go func() {
		for { // check go channel is closed or not
			closed := false
			select {
			case _, ok := <-done:
				if !ok {
					closed = true
				}
			default:
			}

			fmt.Printf("channel is %t\n", closed)
			<-time.After(200 * time.Millisecond)
		}
	}()

	done <- true
	done <- false
	<-time.After(time.Second)

	close(done)

	<-time.After(time.Second)
}
