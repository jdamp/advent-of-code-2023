package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Structures and Datatypes

type Almanac struct{
	Seeds []int;
	Mappings []RangeMapping;
}

func (a Almanac) PassThroughAll(input int) int{
	for _, mapping := range a.Mappings {
		input = mapping.PassThrough(input)
	}
	return input
}

func (a Almanac) IsInSeedRange(input int) bool {
	for i := 0; i < len(a.Seeds)/2; i++ {
		if (input >= a.Seeds[2*i]) && (input < a.Seeds[2*i] + a.Seeds[2*i+1]) {
			return true
		}
	}
	return false
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
		result := al.PassThroughAll(seed)
		if result < locMin {
			locMin = result
		}		
	}
	return locMin
}

func SolvePart2(al Almanac, initialStep int) int {
	// Idea: reverse the mappings, and traverse back starting from location numbers until I obtain a valid seed
	for _, mapping := range al.Mappings {
		for i, mapRange := range mapping.Ranges {
			mapping.Ranges[i] = Range{mapRange.Source, mapRange.Target, mapRange.Length}
		}
	}
	// Now also reverse the order of maps
	slices.Reverse(al.Mappings)

	// To speed up the procedure, run with a initial step size > 1 (e.g. 10000),
	// search for the first location that is mapped to one of the seed ranges.
	// Then use the previous location (i.e. first location - step), and restart
	// the search using a reduced seed size
	startPos := 0
	next := 0
	prev := 0	
	for initialStep >= 1 {
		prev, next = StepPart2(al, startPos, initialStep)
		initialStep = initialStep/10
		startPos = prev
	}
	return next;
	
}

func StepPart2(al Almanac, position, step int) (prev, next int){
	for true {
		result := al.PassThroughAll(position)		
		if al.IsInSeedRange(result){
			return position - step, position
		}
		position += step
	}
	return 0, 0
}


func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(content)
	almanac := ParseInput(input)
	part1 := Solve(almanac)
	part2 := SolvePart2(almanac, 10000)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)

}