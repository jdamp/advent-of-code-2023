package main

import (
	_ "embed"
	"testing"
)

func TestGame(t *testing.T) {
	GameRecords := []struct {
		game       Game
		IsPossible bool
	}{
		{
			IsPossible: true,
			game: Game{
				Id:    1,
				Cubes: []CubeSet{{Red: 4, Green: 0, Blue: 3}, {Red: 1, Green: 2, Blue: 6}, {Red: 0, Green: 2, Blue: 0}},
			},
		},
		{
			IsPossible: false,
			game: Game{
				Id:    3,
				Cubes: []CubeSet{{Red: 20, Green: 8, Blue: 6}, {Red: 4, Green: 13, Blue: 5}, {Red: 1, Green: 5, Blue: 0}},
			},
		},
	}
	for _, GameRecord := range GameRecords {
		got := GameRecord.game.IsValid(CubeSet{Red: 12, Green: 13, Blue: 14})
		if got != GameRecord.IsPossible {
			t.Errorf("Wrong result for game %#v got %v, want %v", GameRecord.game, got, GameRecord.IsPossible)
		}
	}
}

//go:embed test_input.txt
var inputTest string

func TestPart1(t *testing.T) {
	want := 8
	got := SolvePart1(inputTest, CubeSet{Red: 12, Green: 13, Blue: 14})
	if got != want {
		t.Errorf("Wrong result, got %v, want %v", got, want)
	}

}

func TestPart2(t *testing.T) {
	want := 2286
	got := SolvePart2(inputTest)
	if got != want {
		t.Errorf("Wrong result, got %v, want %v", got, want)
	}

}
