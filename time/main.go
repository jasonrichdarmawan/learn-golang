package main

import (
	"fmt"
	"time"
)

func main() {
	dates := []string{"Jan2022", "Feb2022", "Mar2022", "Apr2022", "May2022", "Jun2022", "Jul2022", "Aug2022"}
	for _, date := range dates {
		parsedTime, err := time.Parse("Jan2006", date)
		if err != nil {
			panic(err)
		}
		fmt.Println(parsedTime)
	}
}
