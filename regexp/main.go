package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`(?m)^(?: {2,}(?P<TANGGAL>[0-9]{2}/[0-9]{2}))?(?: {2,21}(?P<KETERANGAN1>[\w/:&.,()-]+(?: [\w/:&.,()-]+)*))?(?: {2,64}(?P<KETERANGAN2>[\w/:&.,()-]+(?: [\w/:&.,()-]+)*))?(?: {2,96}(?P<MUTASI>[\d,.]+(?: (?:DB|CR))*))?(?: {2,}(?P<SALDO>[\d,.]+))?$`)
	for _, subexpname := range regex.SubexpNames() {
		fmt.Println(subexpname, regex.SubexpIndex(subexpname))
	}

	matches := regex.FindAllSubmatch([]byte("     01/07          SALDO AWAL                                                                                                                       1,909,143.49"), -1)

	fmt.Println(string(matches[0][0]))
	for _, match := range matches {
		for _, subexpname := range regex.SubexpNames() {
			if subexpIndex := regex.SubexpIndex(subexpname); subexpIndex != -1 {
				fmt.Println(subexpname, string(match[subexpIndex]))
			}
		}
	}

	fmt.Println(regex.NumSubexp())
}
