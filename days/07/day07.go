package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)


func Solve(input string, joker bool) int64{
	var hands []*Hand
	for _, line := range strings.Split(input, "\n") {
		hands = append(hands, NewHand(line, joker))
	}
	// Sort hands in a way that the weakest ranks occur first
	sort.Slice(hands, func(i, j int) bool {
		if (hands[i].Rank != hands[j].Rank) {
			return hands[i].Rank < hands[j].Rank
		}
		// Compare the cards one-by-one and for the second ordering rule
		for iCard := 0; iCard < len(hands[i].Cards); iCard++ {
			if hands[i].Cards[iCard] < hands[j].Cards[iCard] {
				return true // h[i] < h[j]
			} else if hands[i].Cards[iCard] > hands[j].Cards[iCard] {
				return false // h[i] > h[j]
			}
		}
		return true
	})
	
	var result int64 = 0
	for i, hand := range hands {
		fmt.Println((i+1), hand)
		result += int64((i + 1) * hand.Bid)
	}
	return result
}

func main() {
	content, _ := os.ReadFile("input.txt")
	input := string(content)
	result := Solve(input, false)
	fmt.Printf("Part1: %d\n", result)
	resultPart2 := Solve(input, true)
	fmt.Printf("Part2: %d\n", resultPart2)
}



