package year2015

import _ "embed"

//go:embed inputs/day03.txt
var day03_2015 string

type Coordinates struct {
	X int
	Y int
}

func move(current Coordinates, direction rune) Coordinates {
	switch direction {
	case '>':
		current.X++
	case '<':
		current.X--
	case '^':
		current.Y++
	case 'v':
		current.Y--
	}
	return current
}

func getCoordinatesSet(day03_2015 string, santaCount int) map[Coordinates]struct{} {
	coordinatesSet := make(map[Coordinates]struct{})

	santas := make([]Coordinates, santaCount)
	for i := 0; i < santaCount; i++ {
		santas[i] = Coordinates{X: 0, Y: 0}
	}

	coordinatesSet[santas[0]] = struct{}{}

	for i, direction := range day03_2015 {
		currentSantaIndex := i % santaCount
		santas[currentSantaIndex] = move(santas[currentSantaIndex], direction)
		coordinatesSet[santas[currentSantaIndex]] = struct{}{}
	}

	return coordinatesSet
}

type Day03 struct{}

func (d Day03) SolvePart1() int {
	return len(getCoordinatesSet(day03_2015, 1))
}

func (d Day03) SolvePart2() int {
	return len(getCoordinatesSet(day03_2015, 2))
}

func (d Day03) GetName() string {
	return "Day 03"
}