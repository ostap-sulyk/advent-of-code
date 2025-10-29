package main

import (
	"github.com/ostap-sulyk/advent-of-code/utils"
	"github.com/ostap-sulyk/advent-of-code/year2015"
)

func reverse(s string) string {
	var reversed string
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}
func main() {

	println(utils.GetTree())

	utils.PrintDays(2015, year2015.AllDays())
}
