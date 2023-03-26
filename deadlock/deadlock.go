package main

import (
	"fmt"
	"sync"
)

func main() {
	collection := NewCollection()
	collection.Add("key", "value")
	collection.changeMe()
	fmt.Println(collection.data)
}

type Collection struct {
	mutex sync.Mutex
	data  map[string]string
}

func NewCollection() *Collection {
	return &Collection{
		mutex: sync.Mutex{},
		data:  make(map[string]string),
	}
}

func (c *Collection) Has(key string) bool {
	// c.mutex.Lock() // This is causing the deadlock
	// defer c.mutex.Unlock()
	_, ok := c.data[key]
	return ok
}

func (c *Collection) Add(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Has(key) {
		fmt.Println("Key already exists")
		return
	}
	c.data[key] = value
}

func (c Collection) changeMe() {
	c.data = make(map[string]string)
}
