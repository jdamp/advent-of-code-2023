package main

import (
	"math"

	mapset "github.com/deckarep/golang-set/v2"
)

// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// -> set of numbers for winning numbers,  set of numbers for my numbers
// score: count set intersection, 2**(n-1)

type Card struct {
	MyNumbers      []int
	WinningNumbers []int
}

func (c Card) Score() (score int) {
	mine := mapset.NewSet[int](c.MyNumbers...)
	win := mapset.NewSet[int](c.WinningNumbers...)
	nWin := mine.Intersect(win).Cardinality()
	return int(math.Pow(float64(2), float64(nWin-1)))
}

func main() {

}
