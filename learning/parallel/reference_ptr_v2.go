package parallel

import "sync"

type clientsync struct {
	val int
}

func (c *clientsync) Add(x int) int {
	return c.val + x
}

type librarySync struct {
	c *clientsync
	mx sync.Mutex
}

func NewLibrarySync() *librarySync {
	c := &clientsync{
		val: 0,
	}

	return &librarySync{
		c: c,
	}
}

func (l *librarySync) C() *clientsync {
	l.mx.Lock()
	defer l.mx.Unlock()
	return l.c
}

func (l *librarySync) UpdateClient() {
	l.mx.Lock()
	defer l.mx.Unlock()
	newVal := l.c.val + 1
	l.c = &clientsync{newVal}
}
