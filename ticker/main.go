package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cache := NewCacheWithExpiry()
	cache.addItem("key1", time.Now().Add(time.Second*5))
	cache.addItem("key2", time.Now().Add(time.Second*3))

	time.Sleep(time.Second * 10)
}

type customCacheWithExpiry struct {
	addrs sync.Map
}

func NewCacheWithExpiry() *customCacheWithExpiry {
	return &customCacheWithExpiry{
		addrs: sync.Map{},
	}
}

func (c *customCacheWithExpiry) addItem(key string, val time.Time) {
	fmt.Printf("storing key %s which expires at %s\n", key, val)
	c.addrs.Store(key, val)

	ticker := time.NewTicker(val.Sub(time.Now()))
	go func() {
		<-ticker.C
		fmt.Printf("deleting key %s\n", key)
		c.addrs.Delete(key)
	}()
}
