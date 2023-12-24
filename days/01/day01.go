package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parsePart1(line string) (result int) {
	regex, _ := regexp.Compile("\\d")
	matches := regex.FindAllString(line, -1)
	first := matches[0]
	last := first
	if len(matches) > 1 {
		last = matches[len(matches)-1]
	}
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

func parsePart2(line string) (result int) {
	regex, _ := regexp.Compile("(\\d|one|two|three|four|five|six|seven|eight|nine)")
	matches := regex.FindAllStringSubmatch(line, -1)
	fmt.Println(matches)
	first := translate(matches[0][0])
	last := first
	if len(matches) > 1 {
		last = translate(matches[len(matches)-1][0])
	}
	result, _ = strconv.Atoi(first + last)
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
