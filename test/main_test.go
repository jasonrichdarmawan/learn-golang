package main

import (
	"flag"
	"testing"
)

var loc = flag.String("loc", "", "Name of location to greet")

func TestGreet(t *testing.T) {
	res := Greet(*loc)

	if res != "Hello, San Francisco!" {
		t.Errorf("String mismatch on test")
	}
}
