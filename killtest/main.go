package main

import (
	"fmt"
	"log"
	"time"
)

type Test struct{}

var counter int

func (t *Test) test() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			counter += 1
			log.Println(counter)
		}
	}()
}

func main() {
	t := Test{}
	t.test()
	fmt.Println("this should print immediately")
	time.Sleep(5 * time.Second)
	t = Test{}
	_ = t
}
