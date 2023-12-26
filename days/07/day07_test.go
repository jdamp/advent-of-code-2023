package main

import (
	_ "embed"
	"testing"
)


func TestHandType(t *testing.T) {
	testCases := []struct {
		hand *Hand;
		want rank
	} {
		{NewHand("AAAAA 123"), FiveOfKind},
		{NewHand("AA8AA 28"), FourOfKind},
		{NewHand("23332 34"), FullHouse},
		{NewHand("TTT98 77"), ThreeOfKind},
		{NewHand("23432 998"), TwoPair},
		{NewHand("A23A4 654"), OnePair},
		{NewHand("23456 234"), HighCard},
	}
	for _, testCase := range testCases {
		got := testCase.hand.Rank
		if got != testCase.want {
			t.Errorf("Got %d, want %d for Hand %v", got, testCase.want, testCase.hand)
		}
	}

}

//go:embed test_input.txt
var input string
func TestSolve(t *testing.T) {
	var want int64 = 6440
	got := Solve(input)
	if got != want {
		t.Errorf("Wrong result, got %d, want %d", got, want)
	}

}