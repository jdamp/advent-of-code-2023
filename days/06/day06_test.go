package main

import "testing"

func TestSolvePart1(t *testing.T) {
	times := []int{7, 15, 30}
	distances := []int{9, 40, 200}
	
	want := 288
	got := SolvePart1(times, distances)
	if want != got {
		t.Errorf("Error solving part 1, got %d, want %d", got, want)
	}

	want2 := 71503
	got2 := SolvePart2(71530, 940100)
	if want2 != got2 {
		t.Errorf("Error solving part 1, got %d, want %d", got, want)
	}

}