package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CubeSet struct {
	Red   int
	Blue  int
	Green int
}

func (c CubeSet) IsValid(allCubes CubeSet) bool {
	return c.Red <= allCubes.Red && c.Blue <= allCubes.Blue && c.Green <= allCubes.Green
}

func makeCubeSet(linePart string) (cube CubeSet) {
	cubeRegex := regexp.MustCompile(`(\d+) (\w+)`)
	for _, match := range cubeRegex.FindAllStringSubmatch(linePart, -1) {
		color := match[2]
		value, _ := strconv.Atoi(match[1])
		if color == "red" {
			cube.Red = value
		}
		if color == "blue" {
			cube.Blue = value
		}
		if color == "green" {
			cube.Green = value
		}
	}
	return cube
}

type Game struct {
	Id    int
	Cubes []CubeSet
}

func (g Game) IsValid(allCubes CubeSet) bool {
	valid := true
	for _, cube := range g.Cubes {
		valid = valid && cube.IsValid(allCubes)
	}
	return valid
}

func (g Game) GetMin() CubeSet {
	minCubes := g.Cubes[0]
	for _, cubes := range g.Cubes[1:] {
		minCubes.Red = max(cubes.Red, minCubes.Red)
		minCubes.Blue = max(cubes.Blue, minCubes.Blue)
		minCubes.Green = max(cubes.Green, minCubes.Green)
	}
	return minCubes
}

func MakeGame(line string) (game Game) {
	idRegexp := regexp.MustCompile(`(\d+)`)
	lineParts := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(idRegexp.FindStringSubmatch(lineParts[0])[1])
	game.Id = gameId
	for _, dice := range strings.Split(lineParts[1], ";") {
		game.Cubes = append(game.Cubes, makeCubeSet(dice))
	}
	return game
}

func SolvePart1(input string, allCubes CubeSet) (idSum int) {

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := MakeGame(line)
		if game.IsValid(allCubes) {
			idSum += game.Id
		}
	}
	return idSum
}

func SolvePart2(input string) (power int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := MakeGame(line)
		minCubes := game.GetMin()
		power += minCubes.Blue * minCubes.Red * minCubes.Green
	}
	return power
}

//go:embed input.txt
var input string

func main() {
	allCubes := CubeSet{Red: 12, Green: 13, Blue: 14}
	fmt.Printf("Part 1: %d\n", SolvePart1(input, allCubes))
	fmt.Printf("Part 2: %d\n", SolvePart2(input))
}
