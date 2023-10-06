package main

import (
	"fmt"

	"github.com/wesh92/learning_golang/pkg/httpstuff/aocget" // adjust this import path based on your module name
)

func main() {
	body, err := aocget.AoCRequest("https://adventofcode.com/2020/day/1/input", "session=53616c7465645f5f8f8c62e3a9464a86199f1d08848653e924f0cacb73b0778c086e27eac56006b47dcdbcffa79a788cd30fecaa1c85547a3ee17f0cc731e304")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(body)
}
