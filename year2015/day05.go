package year2015

import (
	_ "embed"
	"strings"
)

//go:embed inputs/day05.txt
var day05_2015 string

type Day05 struct{}

var vovels = "aeiou"
var excludes = []string{"ab", "cd", "pq", "xy"}

func stringIsNice(s string) bool {
	for _, exclude := range excludes {
		if strings.Contains(s, exclude) {
			return false
		}
	}
	hasLetterAppearingTwiceInARow := false
	countVovels := 0
	for i := 0; i < len(s); i++ {
		if strings.Contains(vovels, (string)(s[i])) {
			countVovels++
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			hasLetterAppearingTwiceInARow = true
		}

	}

	return hasLetterAppearingTwiceInARow && countVovels >= 3
}

func (d Day05) SolvePart1() int {
	niceStringsCount := 0
	for _, line := range strings.Split(day05_2015, "\n") {
		if stringIsNice(line) {
			niceStringsCount++
		}
	}
	return niceStringsCount
}

func (d Day05) SolvePart2() int {
	return 0
}

func (d Day05) GetName() string {
	return "Day 05"
}
