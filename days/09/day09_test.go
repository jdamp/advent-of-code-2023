package main

import (
	"testing"
)
func TestSequence(t *testing.T) {
	testCases := []struct {
		values []int
		want int
	} {
		{
			values: []int{0, 3, 6, 9, 12, 15},
			want: 18,
		},
		{
			values: []int{1, 3, 6, 10, 15, 21},
			want: 28,
		},
		{
			values: []int{10, 13, 16, 21, 30, 45},
			want: 68,
		},
	}
	for _, testCase := range testCases {
		y := testCase.values
		n := len(y)
		got := LangrangeInterpolation(y)(n+1)
		if got != float64(testCase.want) {
			t.Errorf("Error, want %d, got %.2f", testCase.want, got)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	lines := []string {
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	want1 := 114.
	want2 := 2.
	got1, got2, _ := SolveBoth(lines)
	if got1 != want1 {
		t.Errorf("Wrong result for Part1, got %.2f, want %.2f", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("Wrong result for Part1, got %.2f, want %.2f", got2, want2)
	}
}