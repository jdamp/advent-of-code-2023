package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parsePart1(line string) (result int) {
	regex := "(\\d)"
	return parse(line, regex)
}

func parsePart2(line string) (result int) {
	regex := "(\\d|one|two|three|four|five|six|seven|eight|nine)"
	return parse(line, regex)
}

func parse(line string, baseExp string) (result int) {
	start_exp, _ := regexp.Compile("^.*?" + baseExp)
	end_exp, _ := regexp.Compile(".*" + baseExp)
	matches_start := start_exp.FindStringSubmatch(line)
	matches_end := end_exp.FindStringSubmatch(line)
	first := translate(matches_start[1])
	last := translate(matches_end[1])
	result, _ = strconv.Atoi(first + last)
	return result
}

func translate(digit string) string {
	dictionary := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6",
		"seven": "7", "eight": "8", "nine": "9",
	}
	result, ok := dictionary[digit]
	if !ok {
		// Key not in map -> return input
		result = digit
	}
	return result
}
func solve(input string, parser func(string) int) (result int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		result += parser(line)
	}
	return result
}

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Day 01 part 1: %d\n", solve(input, parsePart1))
	fmt.Printf("Day 01 part 2: %d\n", solve(input, parsePart2))
}
