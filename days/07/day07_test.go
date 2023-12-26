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
		{NewHand("AAAAA 123", false), FiveOfKind},
		{NewHand("AA8AA 28", false), FourOfKind},
		{NewHand("23332 34", false), FullHouse},
		{NewHand("TTT98 77", false), ThreeOfKind},
		{NewHand("23432 998", false), TwoPair},
		{NewHand("A23A4 654", false), OnePair},
		{NewHand("23456 234", false), HighCard},
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
	got := Solve(input, false)
	if got != want {
		t.Errorf("Wrong result, got %d, want %d", got, want)
	}
	gotPart2 := Solve(input, true)
	var wantPart2 int64 = 5905
	if gotPart2 != wantPart2 {
		t.Errorf("Wrong result, got %d, want %d", gotPart2, wantPart2)
	}
}


func TestReplaceJoker(t *testing.T) {
	testCases := []struct {
		hand *Hand;
		want rank
	} {
		{NewHand("52T3K 765", true), HighCard},
		{NewHand("2345J 28", true), OnePair},
		{NewHand("2233J 28", true), FullHouse},
		{NewHand("QKAJJ 320", true), ThreeOfKind},
		{NewHand("2234J 320", true), ThreeOfKind},
		{NewHand("T55J5 684", true), FourOfKind},
		{NewHand("QQQJJ 77", true), FiveOfKind},
		{NewHand("JJJJJ 1", true), FiveOfKind},
		{NewHand("JJJJQ 2", true), FiveOfKind},
		{NewHand("JJJ66 2", true), FiveOfKind},
	}

	for _, testCase := range testCases {
		if testCase.hand.Rank != testCase.want {
			t.Errorf("Error replacing Jokers: got %d, wanted %d for %v",
			 testCase.hand.Rank, testCase.want, testCase.hand)
		}
	}
}