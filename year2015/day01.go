package year2015

import _ "embed"

//go:embed inputs/day01.txt
var day01_2015 string

type Day01 struct{}

func (d Day01) SolvePart1() int {
	floor := 0
	for i := 0; i < len(day01_2015); i++ {
		if day01_2015[i] == '(' {
			floor++
		} else if day01_2015[i] == ')' {
			floor--
		}
	}

	return floor
}

func (d Day01) SolvePart2() int {
	floor := 0
	position := 0
	for i := 0; i < len(day01_2015); i++ {
		if day01_2015[i] == '(' {
			floor++
		} else if day01_2015[i] == ')' {
			floor--
		}
		if floor == -1 {
			position = i
			break
		}
	}

	return position + 1
}

func (d Day01) GetName() string {
	return "Day 01"
}
