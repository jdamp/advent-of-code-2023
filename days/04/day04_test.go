package main

import (
	_ "embed"
	"testing"
)

func TestCard(t *testing.T) {
	CardTests := []struct {
		card Card
		want int
	}{
		{
			card: Card{
				MyNumbers:      []int{83, 86, 6, 31, 17, 9, 48, 53},
				WinningNumbers: []int{41, 48, 83, 86, 17},
			},
			want: 8,
		},
		{
			card: Card{
				MyNumbers:      []int{61, 30, 68, 82, 17, 32, 24, 19},
				WinningNumbers: []int{13, 32, 20, 16, 61},
			},
			want: 2,
		},
		{
			card: Card{
				MyNumbers:      []int{69, 82, 63, 72, 16, 21, 14, 1},
				WinningNumbers: []int{1, 21, 53, 59, 44},
			},
			want: 2,
		},
		{
			card: Card{
				MyNumbers:      []int{59, 84, 76, 51, 58, 5, 54, 83},
				WinningNumbers: []int{41, 92, 73, 84, 69},
			},
			want: 1,
		},
		{
			card: Card{
				MyNumbers:      []int{88, 30, 70, 12, 93, 22, 82, 36},
				WinningNumbers: []int{87, 83, 26, 28, 32},
			},
			want: 0,
		},
		{
			card: Card{
				MyNumbers:      []int{74, 77, 10, 23, 35, 67, 36, 11},
				WinningNumbers: []int{31, 18, 13, 56, 72},
			},
			want: 0,
		},
	}

	for _, cardTest := range CardTests {
		got := cardTest.card.Score()
		if got != cardTest.want {
			t.Errorf("Wrong result, got %d, want %d", got, cardTest.want)
		}
	}
}


//go:embed test_input.txt
var test_input string
func TestSole(t *testing.T) {
	got1, got2 := Solve(test_input)
	
	want1 := 13
	want2 := 30
	if got1 != want1 {
		t.Errorf("Wrong result for Part1, got %d, want %d", got1, want1)	
	}
	if got2 != want2 {
		t.Errorf("Wrong result for Part2, got %d, want %d", got2, want2)
	}
}


//Part 2:
//-> Need score per card
