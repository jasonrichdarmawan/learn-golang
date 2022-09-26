package main

import (
	"flag"
	"fmt"
	"log"
)

func Greet(loc string) string {
	return "Hello, " + loc + "!"
}

func main() {
	loc := flag.String("loc", "", "Name of location to greet")
	flag.Parse()
	if *loc == "" {
		log.Fatal("variable loc should not be empty")
	}

	fmt.Println(Greet(*loc))
}
