package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	Counter = 0

	start := time.Now()

	addrs := map[int]time.Time{}

	for i := 0; i < 1500; i++ {
		addrs[i] = time.Now()
	}

	if len(addrs) != 1500 {
		t.Fatalf("race condition")
	}

	log.Println("map[int]time.Time", time.Since(start))
}

type AddrsMutex struct {
	sync.Mutex
	addrs map[int]time.Time
}

var Counter int

func put(addrs *AddrsMutex, wg *sync.WaitGroup) {
	addrs.Lock()
	for i := 0; i < 500; i++ {
		Counter += 1
		addrs.addrs[Counter] = time.Now()
	}
	addrs.Unlock()
	wg.Done()
}

func TestMutex(t *testing.T) {
	start := time.Now()

	addrs := AddrsMutex{sync.Mutex{}, map[int]time.Time{}}

	var wg sync.WaitGroup

	wg.Add(3)
	go put(&addrs, &wg)
	go put(&addrs, &wg)
	go put(&addrs, &wg)

	wg.Wait()

	if len(addrs.addrs) != 1500 {
		t.Fatal("race condition")
	}

	log.Println("Sync.Mutex", time.Since(start))
}

func putMap(addrs *sync.Map, wg *sync.WaitGroup, startFrom int) {
	for i := 0; i < 500; i++ {
		startFrom += 1
		addrs.Store(startFrom, time.Now())
	}
	wg.Done()
}

func TestSyncMap(t *testing.T) {
	Counter = 0

	start := time.Now()

	addrs := sync.Map{}

	var wg sync.WaitGroup

	wg.Add(3)
	go putMap(&addrs, &wg, 0)
	go putMap(&addrs, &wg, 501)
	go putMap(&addrs, &wg, 1001)

	wg.Wait()

	log.Println("Sync.Map ", time.Since(start))

	var count int

	addrs.Range(func(key, value any) bool {
		count += 1
		return true
	})

	if count != 1500 {
		log.Fatal("race condition, expected len: 1500 actual: ", count)
	}
}
