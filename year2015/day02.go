package year2015

import (
	"sort"
	"strconv"
	"strings"
	_ "embed"
)

//go:embed inputs/day02.txt
var day02_2015 string


type Day02 struct{}

type Prism struct {
	W int
	H int
	L int
}

func (p *Prism) getSurface() int {
	return (2 * p.L * p.W) + (2 * p.W * p.H) + (2 * p.H * p.L)
}

func (p *Prism) getSmallestArea() int {
	areas := []int{p.L * p.W, p.W * p.H, p.H * p.L}
	sort.Ints(areas)
	return areas[0]
}

func (p *Prism) getRibbon() int {
	dims := []int{p.W, p.H, p.L}
	sort.Ints(dims)
	return 2*(dims[0]+dims[1]) + p.W*p.H*p.L
}

func createPrism(dimensions string) Prism {
	splitted := strings.Split(dimensions, "x")

	w, _ := strconv.Atoi(splitted[0])
	h, _ := strconv.Atoi(splitted[1])
	l, _ := strconv.Atoi(splitted[2])

	return Prism{W: w, H: h, L: l}
}

func parseDimensions(input string) []Prism {
	dimensions := strings.Split(input, "\n")
	var prisms []Prism

	for _, dimension := range dimensions {
		prisms = append(prisms, createPrism(dimension))
	}

	return prisms
}

func (d Day02) SolvePart1() int {
	prisms := parseDimensions(day02_2015)

	totalSurface := 0
	for _, prism := range prisms {
		totalSurface += prism.getSurface() + prism.getSmallestArea()
	}
	return totalSurface
}

func (d Day02) SolvePart2() int {
	prisms := parseDimensions(day02_2015)

	totalRibbon := 0
	for _, prism := range prisms {
		totalRibbon += prism.getRibbon()
	}
	return totalRibbon
}

func (d Day02) GetName() string {
	return "Day 02"
}
