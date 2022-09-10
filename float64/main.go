package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	mutasis := [][]byte{[]byte("28,000.00 DB"), []byte("5,000,000.00")}
	re := regexp.MustCompile(`(?m)^(?P<MUTASI>[\d,.]+)?(?: (?P<ENTRY>DB))?$`)
	matches := re.FindAllSubmatch(bytes.Join(mutasis, []byte("\n")), -1)
	for _, match := range matches {
		fmt.Println(string(match[1]))
	}
}
