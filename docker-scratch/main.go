package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Hello World!", os.Getuid(), os.Getgid())
}
