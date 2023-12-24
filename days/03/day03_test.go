package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	wantSym := Symbols{
		{3, 1}: "*",
		{6, 3}: "#",
		{3, 4}: "*",
		{5, 5}: "+",
		{3, 8}: "$",
		{5, 8}: "*",
	}
	wantNum := Numbers{
		{0, 0}: {467, 3},
		{5, 0}: {114, 3},
		{2, 2}: {35, 2},
		{6, 2}: {633, 3},
		{0, 4}: {617, 3},
		{7, 5}: {58, 2},
		{2, 6}: {592, 3},
		{6, 7}: {755, 3},
		{1, 9}: {664, 3},
		{5, 9}: {598, 3},
	}

	gotSym, gotNum := Parse("test_input.txt")
	if !reflect.DeepEqual(gotSym, wantSym) {
		t.Errorf("Wrong symbols parsed, \ngot %+v, \nwant %+v", gotSym, wantSym)
	}

	if !reflect.DeepEqual(gotNum, wantNum) {
		t.Errorf("Wrong numbers parsed, \ngot %+v, \nwant %+v", gotNum, wantNum)
	}
}

func TestSolvePart1(t *testing.T) {
	sym, num := Parse("test_input.txt")
	got := SolvePart1(sym, num)
	want := 4361
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestSolvePar2(t *testing.T) {
	sym, num := Parse("test_input.txt")
	got := SolvePart2(sym, num)
	want := 467835
	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
