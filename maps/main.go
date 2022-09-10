// reference: https://gobyexample.com/maps

package main

import "fmt"

type taccount struct {
	DATE        string
	DESCRIPTION string
	AMOUNT      int
}

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys
	// and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The for...range loop statement can be used to fetch the key and value of a map.
	for key, value := range n {
		fmt.Println("key:", key, "value:", value)
	}

	// []interface{} can be used to make a slice with multiple types
	taccountsinterface := map[string][][]interface{}{"FAMILY MART": {{"01/07", "0107/TEEEZ/EZ9E313", 28000}}}
	fmt.Println(`taccountsinterface["FAMILY MART"][0]`, taccountsinterface["FAMILY MART"][0])

	taccounttruct := taccount{"02/07", "0207/TEEEZ/EZ9E313", 30000}
	fmt.Println(`taccounttruct.DATE`, taccounttruct.DATE)

	taccountstruct := map[string][]taccount{
		"FAMILY MART": {
			{"01/07", "0107/TEEEZ/EZ9E313", 28000},
			taccounttruct,
		},
	}

	fmt.Println(`taccountstruct["FAMILY MART"][0].DATE`, taccountstruct["FAMILY MART"][0].DATE)
}
