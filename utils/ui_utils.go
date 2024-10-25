package utils

import (
	"fmt"
)

func PrintDays(year int, days []Day) {
	fmt.Printf("Year %d:\n", year)
	for _, day := range days {
		fmt.Println(Format(day))
	}
	fmt.Println()
}
