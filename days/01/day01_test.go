package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var inputTest string

//go:embed test_input2.txt
var inputTest2 string

func TestParseLine(t *testing.T) {
	lineTests := []struct {
		line  string
		want1 int
		want2 int
	}{
		{"1abc2", 12, 12},
		{"df6jfn3five", 63, 65},
		{"14resd56t7", 17, 17},
		{"a2seven", 22, 27},
		{"1oneight", 11, 18},
	}
	for _, tt := range lineTests {
		got1 := parsePart1(tt.line)
		if got1 != tt.want1 {
			t.Errorf("got %d want %d", got1, tt.want1)
		}
		got2 := parsePart2(tt.line)
		if got2 != tt.want2 {
			t.Errorf("got %d want %d", got2, tt.want2)
		}
	}
}

func TestPart1(t *testing.T) {
	result := solve(inputTest, parsePart1)
	expected := 142
	if result != expected {
		t.Errorf("Wrong result, got %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := solve(inputTest2, parsePart2)
	expected := 281
	if result != expected {
		t.Errorf("Wrong result, got %d, want %d", result, expected)
	}
}
