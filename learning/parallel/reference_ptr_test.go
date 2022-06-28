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
		go func() {
			l.UpdateClient()
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestRaceTest2(t *testing.T) {
	l := parallel.NewLibrarySync()

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
		go func() {
			l.UpdateClient()
			wg.Done()
		}()
	}

	wg.Wait()
}
