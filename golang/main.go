package main

import (
	"fmt"
	"github.com/jmontroy90/exercism-4824/week1"
)

func main() {
	years := []int{1997, 1900, 2000, 1996}
	for _, yr := range years {
		fmt.Printf("%v is leap year: %v\n", yr, week1.IsLeapYear(yr))
	}
}
