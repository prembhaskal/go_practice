package parallel

import "sync"

type client struct {
	val int
}

func (c *client) Add(x int) int {
	return c.val + x
}

type library struct {
	c *client
	mx sync.Mutex
}

func NewLibrary() *library {
	c := &client{
		val: 0,
	}

	return &library{
		c: c,
	}
}

func (l *library) C() *client {
	l.mx.Lock()
	defer l.mx.Unlock()
	return l.c
}

func (l *library) UpdateClient() {
	l.mx.Lock()
	defer l.mx.Unlock()
	newVal := l.c.val + 1
	l.c = &client{newVal}
}
