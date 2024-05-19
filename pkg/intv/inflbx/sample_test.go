package main

import (
	"testing"
	"fmt"
)

func TestMap(t *testing.T) {

	// test custom map
	cm := newCmap()
	cm.Put("k1", "value1")
	cm.Put("k2", "value2")

	k1V, present := cm.Get("k1")
	fmt.Printf("key: %s, value: %s\n", k1V, present)
	
	k2V, present := cm.Get("k2")
	fmt.Printf("key: %s, value: %s\n", k2V, present)
	
	k3V, present := cm.Get("k3")
	fmt.Printf("key: %s, value: %s\n", k3V, present)
}