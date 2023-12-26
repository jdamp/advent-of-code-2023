package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type rank int

// HandRank
const (
	HighCard rank =iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var OnePairCount = []int{1, 1, 1, 2}
var TwoPairCount = []int{1, 2, 2}
var ThreeOfKindCount = []int{1, 1, 3}
var FullHouseCount = []int{2, 3}
var FourOfKindCount = []int{1, 4}
var FiveOfKindCount = []int{5}

// translation for non-numeric card values
var cardValues = map[string]int {
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
}

var cardValuesJoker = map[string]int {
	"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10,
}


type Hand struct {
	Cards []int
	Rank rank
	Bid int
}

func NewHand(line string, joker bool) *Hand {
	parts := strings.Split(line, " ")
	var hand Hand
	var lookup map[string]int
	if joker {
		lookup = cardValuesJoker
	} else {
		lookup = cardValues
	}
	for _, cardSymbol := range parts[0] {
		if cardValue, err := strconv.Atoi(string(cardSymbol)); err == nil {
			hand.Cards = append(hand.Cards, cardValue)
		} else {
			cardValue = lookup[string(cardSymbol)]
			hand.Cards = append(hand.Cards, cardValue)
		}
	}

	hand.Bid, _ = strconv.Atoi(parts[1])
	hand.Rank = rankHand(hand)

	return &hand
}

func rankHand(hand Hand) rank{
	// Count how many of each card the hand has
	counts := make(map[int]int)
	for _, card := range hand.Cards {
		if _, ok := counts[card]; ok {
			counts[card] += 1
		} else {
			counts[card] = 1
		}		
	}
	// Convert count maps to sorted slice of values, only storing
	// the counts of cards that the hand actually has.
	countValues := []int{}
	for _, value := range counts {
		if value > 0 {
			countValues = append(countValues, value)
		}
	}	
	slices.Sort(countValues)
	if slices.Equal(countValues, FiveOfKindCount) {
		return FiveOfKind
	} else if slices.Equal(countValues, FourOfKindCount) {
		return FourOfKind
	} else if slices.Equal(countValues, FullHouseCount) {
		return FullHouse
	} else if slices.Equal(countValues, ThreeOfKindCount) {
		return ThreeOfKind
	} else if slices.Equal(countValues, TwoPairCount) {
		return TwoPair
	} else if slices.Equal(countValues, OnePairCount) {
		return OnePair
	} else {
		return HighCard
	}
}



func Solve(input string) int64{
	var hands []*Hand
	for _, line := range strings.Split(input, "\n") {
		hands = append(hands, NewHand(line, false))
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
		result += int64((i + 1) * hand.Bid)
	}
	return result
}

func main() {
	content, _ := os.ReadFile("input.txt")
	input := string(content)
	result := Solve(input)
	fmt.Printf("Part1: %d\n", result)

}



// Essentially I just need to recalculate the rank
func GetJokerCardRank(hand Hand, counts map[int]int) rank {
	// Get numbers of jokers from the map
	joker := cardValuesJoker["J"]
	jokerCounts := counts[joker]
	delete(counts, joker)
	// Convert the remaining card to a count array
	countValues := []int{}
	for _, value := range counts {
		if value > 0 {
			countValues = append(countValues, value)
		}
	}
	slices.Sort(countValues)
	switch jokerCounts {
	// 4 or 5 J -> FiveOfKind
	case 5:
		return FiveOfKind
	case 4:
		return FiveOfKind
	// 3 J -> if 2 equal Five, else Four
	case 3:
		if slices.Contains(countValues, 2) {
			return FiveOfKind
		}
		return FourOfKind
	// 2 J -> if 3 equal Five, else if 2 equal Four, else Three
	case 2:
		if slices.Contains(countValues, 3) {
			return FiveOfKind
		} else if slices.Contains(countValues, 2) {
			return FourOfKind
		}
		return ThreeOfKind
	// 1 J -> if 4 equal Five, else if 3 equal Four, else if 2&2 equal full house
	//   else if 2 equal Three, else OnePair
	case 1:
		if slices.Contains(countValues, 4) {
			return FiveOfKind
		} else if slices.Contains(countValues, 3) {
			return FourOfKind
		} else if slices.Equal(countValues, []int{2, 2}) {
			return FullHouse
		} else if slices.Contains(countValues, 2) {
			return ThreeOfKind
		}
		return OnePair
	}
	return hand.Rank
}
