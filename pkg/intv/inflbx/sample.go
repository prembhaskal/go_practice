package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello %s\n", "world")


	// test custom map
	cm := newCmap()
	cm.Put("k1", "value1")
	cm.Put("k2", "value2")

	k1V, present := cm.Get("k1")
	fmt.Printf("key: %s, value: %t\n", k1V, present)
	
	k2V, present := cm.Get("k2")
	fmt.Printf("key: %s, value: %t\n", k2V, present)
	
	k3V, present := cm.Get("k3")
	fmt.Printf("key: %s, value: %t\n", k3V, present)

	fmt.Println("deleting k2") 
	cm.Delete("k2")
	k1V, present = cm.Get("k1")
	fmt.Printf("key: %s, value: %t\n", k1V, present)
	
	k2V, present = cm.Get("k2")
	fmt.Printf("key: %s, value: %t\n", k2V, present)
	
}


// implement a custom map
// put key, get key, delete key

// string vs string
type KeyValuePair struct { key string
	value string
}

type cmap struct {
	data [][]KeyValuePair
	bucketCount int // init to fix value say 32
}

func newCmap() *cmap {
	return &cmap{
		data : make([][]KeyValuePair, 32),
		bucketCount: 32,
	}
}


func (c *cmap) Put(key, value string) {
	// get key 
	// compute a hash
	keyHash := c.hash(key)

	// compute index for the hash
	bktIndex := keyHash % c.bucketCount

	bktDatas := c.data[bktIndex]
	
	// push data into that
	// check if existing data present there
	found := false
	for idx, keypair := range bktDatas {
		if keypair.key == key {
			found = true
			bktDatas[idx] = KeyValuePair{key: key, value: value}
			break
		}
	}

	if !found {
		bktDatas = append(bktDatas, KeyValuePair{key: key, value: value})
	}

	c.data[bktIndex] = bktDatas
}

func (c *cmap) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash += 31 * int(ch)
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (c *cmap) Get(key string) (string, bool) {
	// get key 
	// compute a hash
	keyHash := c.hash(key)

	// compute index for the hash
	bktIndex := keyHash % c.bucketCount

	bktDatas := c.data[bktIndex]
	
	for _, keypair := range bktDatas {
		if keypair.key == key {
			return keypair.value, true
		}
	}

	return "", false
}

func (c *cmap) Delete(key string) {
	// get key 
	// compute a hash
	keyHash := c.hash(key)

	// compute index for the hash
	bktIndex := keyHash % c.bucketCount

	bktDatas := c.data[bktIndex]
	
	for idx, keypair := range bktDatas {
		if keypair.key == key {
			bktDatas[idx] = KeyValuePair{}
			break
		}
	}

}

