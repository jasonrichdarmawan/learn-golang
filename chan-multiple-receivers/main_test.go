package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

type MapMux struct {
	sync.RWMutex
	m map[int]bool
	sync.WaitGroup
}

// careful, this is a hack.
// because mux.Lock(), no goroutine will do a read and write to mux.m at the same time.
// this will fail, if you create another function that read from mux.m and run it with goroutine.
func (mux *MapMux) AddInt(i int) {
	mux.Lock()
	if _, isExist := mux.m[i]; isExist {
		log.Fatal("race condition")
	}
	mux.m[i] = true
	mux.Unlock()
	mux.Done()
}

func TestChanRaceCondition(t *testing.T) {
	log.SetFlags(0)
	chanSize := 1500
	c := make(chan int, chanSize)
	defer close(c)

	writerSize := 100
	for i := 0; i < writerSize; i++ {
		go func(key int) {
			for y := chanSize / writerSize * key; y < chanSize/writerSize*(key+1); y++ {
				time.Sleep(10 * time.Millisecond)
				c <- y
			}
		}(i)
	}

	mux := MapMux{sync.RWMutex{}, map[int]bool{}, sync.WaitGroup{}}

	mux.Add(chanSize)

	readerSize := 100
	for i := 0; i < readerSize; i++ {
		go func(key int) {
			for {
				select {
				case payload, ok := <-c:
					if !ok {
						log.Printf("go%d: chan is closed", key)
						return
					}
					log.Printf("go%d: %d", key, payload)
					mux.AddInt(payload)
				case <-time.After(1 * time.Second):
					// if after 1s there is no new data from the channel
					// just kill the goroutine to save resources.
					log.Printf("go%d: empty channel for more than 1s, gracefully kill this goroutine", key)
					return // this return statement will break the for statement.
				}
			}
		}(i)
	}

	mux.Wait()

	if len(mux.m) != chanSize {
		t.Fatal("expected len: ", chanSize, ", actual len: ", len(mux.m))
	}
}
