package main

import "testing"

func TestCountFindSteps(t *testing.T) {
	testCases := []struct{
		path string;
		want int;
	} {
		{"test_input.txt", 2},
		{"test_input2.txt", 6},
	}
	for _, testCase := range testCases {
		instructions, network := Parse(testCase.path)
		got, _ := network.CountFindSteps("AAA", instructions, LoopCondiditionPart1)
		if got != testCase.want {
			t.Errorf("Error, got %d, want %d", got, testCase.want)
		}
	}
}


func TestCountFindStepsPart2(t *testing.T) {
	instructions, network := Parse("test_input3.txt")
	want := 6
	got := SolvePart2(*network, instructions)
	if got != want {
		t.Errorf("Error, got %d, want %d", got, want)
	}
}