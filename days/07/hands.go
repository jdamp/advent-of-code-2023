package main

import (
	"slices"
	"strconv"
	"strings"
)


type Hand struct {
	Cards []int
	name string
	Rank rank
	Bid int
}

// Build a new Hand
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
	if joker {
		hand.Rank = rankHandJoker(hand)
	} else {
		hand.Rank = rankHand(hand)
	}
	hand.name = parts[0]
	return &hand
}



func rankHand(hand Hand) rank{
	// Count how many of each card the hand has
	counts := countCards(hand.Cards)
	// Convert count maps to sorted slice of values, only storing
	// the counts of cards that the hand actually has.
	countValues := getSortedCountValues(counts)
	
	if slices.Equal(countValues, fiveOfKindCount) {
		return FiveOfKind
	} else if slices.Equal(countValues, fourOfKindCount) {
		return FourOfKind
	} else if slices.Equal(countValues, fullHouseCount) {
		return FullHouse
	} else if slices.Equal(countValues, threeOfKindCount) {
		return ThreeOfKind
	} else if slices.Equal(countValues, twoPairCount) {
		return TwoPair
	} else if slices.Equal(countValues, onePairCount) {
		return OnePair
	} else {
		return HighCard
	}
}


// Essentially I just need to recalculate the rank
func rankHandJoker(hand Hand) rank {
	counts := countCards(hand.Cards)
	// Get numbers of jokers from the map
	joker := cardValuesJoker["J"]
	jokerCounts := counts[joker]
	delete(counts, joker)
	// Convert the remaining cards to a count array
	countValues := getSortedCountValues(counts)
	// Depending on the numbers of jokers, different best cases exist.
	// The actual values of the cards does not matter, only how many of the same cards 
	// a hand as
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
	// No jokers -> use normal ranking from Part 1
	return rankHand(hand)
}


func countCards(cards []int) map[int]int{
	counts := make(map[int]int)
	for _, card := range cards {
		if _, ok := counts[card]; ok {
			counts[card] += 1
		} else {
			counts[card] = 1
		}		
	}
	return counts
}


func getSortedCountValues(counts map[int]int) []int {
	countValues := []int{}
	for _, value := range counts {
		if value > 0 {
			countValues = append(countValues, value)
		}
	}
	slices.Sort(countValues)
	return countValues
}

type rank int

// Different possible ranks
const (
	HighCard rank =iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

// signatures of card counts for the different ranks
var onePairCount = []int{1, 1, 1, 2}
var twoPairCount = []int{1, 2, 2}
var threeOfKindCount = []int{1, 1, 3}
var fullHouseCount = []int{2, 3}
var fourOfKindCount = []int{1, 4}
var fiveOfKindCount = []int{5}

// translation for non-numeric card values
var cardValues = map[string]int {
	"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
}

// the same, but with "J" as the Joker
var cardValuesJoker = map[string]int {
	"A": 14, "K": 13, "Q": 12, "J": 1, "T": 10,
}
