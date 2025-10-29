package utils

import (
	"fmt"
	"strings"
)

const (
	ColorReset     = "\033[0m"
	ColorGreen     = "\033[32m"
	ColorDarkGreen = "\033[38;2;25;50;50m"
	ColorYellow    = "\033[33m"
	ColorBrown     = "\033[38;5;137m"
	ColorRed       = "\033[31m"
	ColorBlue      = "\033[34m"
	ColorWhite     = "\033[37m"
)


const TreeTop = `
.  .    .  .   .  .
	 *   .   .
	 ^
 .    . /|\    .   .
   .   /*|O\
.   . /*/|\*\ .  .
     /X/O|*\X\ .   .
 .  /*/X/|\X\*\  .
   /O/*/X|*\O\X\   .
. /*/O/X/|\X\O\*\.
 /X/O/*/X|O\X\*\O\.
/O/X/*/O/|\X\*\O\X\`

const TreeBottom = `
        |X|
........|X|........`


func Colorize(text string, colorCode string) string {
	return colorCode + text + ColorReset
}

func GetTree() string {

	var newTree string
	for i := range len(TreeTop) {
		switch TreeTop[i] {
		case '*':
			newTree += Colorize("*", ColorYellow)
		case '/', '\\', '|', '^':
			newTree += Colorize(string(TreeTop[i]), ColorDarkGreen)
		case 'X':
			newTree += Colorize("X", ColorRed)
		case 'O':
			newTree += Colorize("O", ColorBlue)
		case '[', ']':
			newTree += Colorize(string(TreeTop[i]), ColorBrown)
		case '.':
			newTree += Colorize(string(TreeTop[i]), ColorWhite)
		default:
			newTree += string(TreeTop[i])
		}
	}
	tree := newTree + strings.ReplaceAll(TreeBottom, "|X|", Colorize("|X|", ColorBrown))

	return tree
}

func PrintDays(year int, days []Day) {
	fmt.Printf("Year %d:\n", year)
	for _, day := range days {
		fmt.Println(Format(day))
	}
	fmt.Println()
}
