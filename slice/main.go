package main

import (
	"fmt"
	"sort"
	"time"
)

type Accounts struct {
	transactions [][]string
}

type ByDate []string

func (v ByDate) Len() int      { return len(v) }
func (v ByDate) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v ByDate) Less(i, j int) bool {
	date1, err := time.Parse("Jan2006", v[i])
	if err != nil {
		panic(err)
	}
	date2, err := time.Parse("Jan2006", v[j])
	if err != nil {
		panic(err)
	}
	return date1.Before(date2)
}

func main() {
	accounts := Accounts{}

	accounts.transactions = append(accounts.transactions, []string{})
	accounts.transactions[0] = append(accounts.transactions[0], "1")
	// fmt.Println(accounts.transactions[0][0])

	// sort
	slices := []string{"Jul2022", "Jun2022", "May2022"}
	sort.Sort(ByDate(slices))
	fmt.Println(slices)
}
