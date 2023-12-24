package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x int
	y int
}

type PartNumber struct {
	value int
	len   int
}

type Symbols map[Position]string
type Numbers map[Position]PartNumber
type GearNumbers map[Position][]int

func Parse(path string) (Symbols, Numbers) {
	symbols := make(Symbols)
	numbers := make(Numbers)
	numPattern := regexp.MustCompile(`\d+`)
	symPattern := regexp.MustCompile(`[^\w.]`)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return symbols, numbers
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		numMatches := numPattern.FindAllStringIndex(line, -1)
		for _, match := range numMatches {
			start, end := match[0], match[1]
			number, _ := strconv.Atoi(line[start:end])
			numbers[Position{start, y}] = PartNumber{number, end - start}
		}
		symMatches := symPattern.FindAllStringIndex(line, -1)
		for _, match := range symMatches {
			start, end := match[0], match[1]
			symbols[Position{start, y}] = line[start:end]
		}

		y++
	}
	return symbols, numbers
}

func SolvePart1(symbols Symbols, numbers Numbers) (sum int) {
	for pos, partNum := range numbers {
		isPartNum := false
		// Check positions in row above and below the number
		for x := pos.x - 1; x <= pos.x+partNum.len; x++ {
			_, okAbove := symbols[Position{x, pos.y - 1}]
			_, okBelow := symbols[Position{x, pos.y + 1}]
			isPartNum = isPartNum || okAbove || okBelow
		}
		// Check two neighboring positions in the same line
		_, okLeft := symbols[Position{pos.x - 1, pos.y}]
		_, okRight := symbols[Position{pos.x + partNum.len, pos.y}]
		isPartNum = isPartNum || okLeft || okRight
		if isPartNum {
			sum += partNum.value
		}
	}
	return sum
}

func checkAddMap(key Position, value int, gears GearNumbers) {
	// Check if the key exists
	if _, ok := gears[key]; ok {
		// Key exists, append to existing slice
		gears[key] = append(gears[key], value)
	} else {
		// Key does not exist, create new slice with the value
		gears[key] = []int{value}
	}
}

func SolvePart2(symbols Symbols, numbers Numbers) (gearSum int) {
	// Step 1: Filter pseudoGears from all symbols
	pseudoGears := make(Symbols)
	for pos, symbol := range symbols {
		if symbol == "*" {
			pseudoGears[pos] = symbol
		}
	}
	// Step 2: Store all neighboring numbers for every gear
	gears := make(GearNumbers)
	for pos, partNum := range numbers {
		// Check positions in row above and below the number
		for x := pos.x - 1; x <= pos.x+partNum.len; x++ {
			above := Position{x, pos.y - 1}
			if _, okAbove := pseudoGears[above]; okAbove {
				checkAddMap(above, partNum.value, gears)
			}
			below := Position{x, pos.y + 1}
			if _, okBelow := pseudoGears[below]; okBelow {
				checkAddMap(below, partNum.value, gears)
			}

		}
		// Check two neighboring positions in the same line
		left := Position{pos.x - 1, pos.y}
		if _, okAbove := pseudoGears[left]; okAbove {
			checkAddMap(left, partNum.value, gears)
		}
		right := Position{pos.x + partNum.len, pos.y}
		if _, okBelow := pseudoGears[right]; okBelow {
			checkAddMap(right, partNum.value, gears)
		}

	}
	for _, gearNums := range gears {
		if len(gearNums) == 2 {
			gearSum += gearNums[0] * gearNums[1]
		}
	}
	return gearSum
}

func main() {
	symbols, numbers := Parse("input.txt")
	fmt.Printf("Part 1: %d\n", SolvePart1(symbols, numbers))
	fmt.Printf("Part 2: %d\n", SolvePart2(symbols, numbers))
}
