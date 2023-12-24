package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Structures and Datatypes

type Almanac struct{
	Seeds []int;
	Mappings []RangeMapping;
}

func (a Almanac) PassThrough() {

}

type Range struct {
	Target int;
	Source int;
	Length int;
}

func (r Range) PassThrough (input int) (int, error) {
	// Assume an input is in exactly one range
	if input >= r.Source && input < r.Source+r.Length {
		return r.Target + (input - r.Source), nil
	}
	return input, errors.New("Not in range!");
}

type RangeMapping struct {
	Ranges []Range;
}

func (rMap RangeMapping) PassThrough (input int) int {
	for _, r := range rMap.Ranges {
		result, err := r.PassThrough(input)
		if err == nil {
			return result
		}
	}
	// All ranges returned errors, which means that we simply return the input
	return input
}

func MakeRangeMapping(input string) RangeMapping{
	var ranges []Range

	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, " ")
		target, _ := strconv.Atoi(numbers[0])
		source, _ := strconv.Atoi(numbers[1])
		length, _ := strconv.Atoi(numbers[2])
		
		ranges = append(ranges, Range{Target: target, Source: source, Length: length})
	}
	return RangeMapping {ranges}
}

// Input parsing
func ParseInput(input string) (alm Almanac){

	firstLine := strings.Split(input, "\n")[0]

	// First line : seeds	
	numExp := regexp.MustCompile(`(\d+)`)
	matches := numExp.FindAllString(firstLine, -1)
	for _, match := range matches {
		seed, err := strconv.Atoi(match)
		if err == nil {
			alm.Seeds = append(alm.Seeds, seed)
		} else {
			log.Fatal(err)
		}
	}
	patterns := []string {
		"seed-to-soil map",
		"soil-to-fertilizer map",
		"fertilizer-to-water map",
		"water-to-light map",
		"light-to-temperature map",
		"temperature-to-humidity map",
		"humidity-to-location map",
	}
	for _, pattern := range patterns {
		alm.Mappings = append(alm.Mappings, MakeRangeMapping(findTextBetween(input, pattern)))
	}
	return alm
}

func findTextBetween(text, phrase string) (string) {
	pattern := fmt.Sprintf(`(?m)^%s:\n(.*(?:\n.+)*?)\n\n`, regexp.QuoteMeta(phrase))
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
	}
	return fmt.Sprintf("No match found for %s", phrase)
}

// Solution:
func Solve(al Almanac) int {
	locMin := math.MaxInt64
	for _, seed := range al.Seeds {
		result := seed
		for _, mapping := range al.Mappings {
			result = mapping.PassThrough(result)
		}
		if result < locMin {
			locMin = result
		}		
	}
	return locMin
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	almanac := ParseInput(input)
	part1 := Solve(almanac)

	fmt.Printf("Part1: %d\n", part1)

}