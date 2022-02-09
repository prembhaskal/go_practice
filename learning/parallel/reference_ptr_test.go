package parallel_test

import (
	"sync"
	"testing"

	"github.com/prembhaskal/go_practice/learning/parallel"
)

func TestRaceTest1(t *testing.T) {
	l := parallel.NewLibrary()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		j := i
		wg.Add(1)
		go func() {
			c := l.C()
			c.Add(j)
			wg.Done()
		}()

		wg.Add(1)
		go func(){
			l.UpdateClient()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			c := l.C()
			for j:=0; j<100;j++ {
				c.Add(1)
			}
		}()
	}

	wg.Wait()
}
