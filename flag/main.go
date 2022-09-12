package main

import (
	"flag"
	"log"
)

var (
	email    = flag.String("e", "", "This is used by CAs, such as Let's Encrypt, to notify about problems with issued certificates.")
	hostname = flag.String("h", "", "This is used by autocert, controls which domains the Manager will attempt to retrieve new certificates for.")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("email:", *email, "hostname:", *hostname)
	flag.Usage()
}
