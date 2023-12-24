package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// -> set of numbers for winning numbers,  set of numbers for my numbers
// score: count set intersection, 2**(n-1)

type Card struct {
	MyNumbers      []int
	WinningNumbers []int
}

func (c Card) CountWins() int {
	mine := mapset.NewSet[int](c.MyNumbers...)
	win := mapset.NewSet[int](c.WinningNumbers...)
	count := mine.Intersect(win).Cardinality()
	return count
}

func (c Card) Score() int {
	nWin := c.CountWins()
	return int(math.Pow(float64(2), float64(nWin-1)))
}
func parseNumbers(s string) []int {
    fields := strings.Fields(s)
    numbers := make([]int, 0, len(fields))

    for _, field := range fields {
        // Skip non-numeric parts (like "Card 1:")
        if num, err := strconv.Atoi(field); err == nil {
            numbers = append(numbers, num)
        }
    }

    return numbers
}

func parseLine(line string) Card {
	parts := strings.Split(line, "|")

	winningNumbers := parseNumbers((strings.TrimSpace(parts[0])))
	myNumbers := parseNumbers((strings.TrimSpace(parts[1])))

	return Card{
		WinningNumbers: winningNumbers,
		MyNumbers: myNumbers,
	}
}

func Solve(input string) (sol01, sol02 int){
	lines := strings.Split(input, "\n")
	// Store how many of which card I currently own, starting with one of each card
	cardCounts := make([]int, len(lines))
	for i := range cardCounts {
		cardCounts[i] = 1
	}

	for iCard, line := range lines {
		score := parseLine(line).Score()		
		sol01 += score
		// Part2: How many matching numbers do I have?
		wins := parseLine(line).CountWins()
		for delta := 1; delta <= wins; delta++ {
			newIndex := iCard+delta
			if newIndex >= len(cardCounts) {
				break
			}			
			cardCounts[newIndex] += cardCounts[iCard]
		}
	}
	// Part 2: Count how many cards I have in total
	for _, cardCount := range cardCounts {
		sol02 += cardCount
	}
	return sol01, sol02
}

//go:embed input.txt
var input string
func main() {
	sol01, sol02 := Solve(input)
	fmt.Printf("Part01: %d\n", sol01)
	fmt.Printf("Part02: %d\n", sol02)
}
