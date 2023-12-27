package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The rows always seem to follow a pattern
//0 0 0 0 0 0 f(x) = 0 | constant 0
//c c c c c c f(x) = a | constant
//            f(x) = a + b*x | linear
//            f(x) = a + b*x + c*x**2 | quadratic
//.....

// Idea: find a polynomial going through all listed points und use this to calculate
// the next value
// Interpolation theorem: For any n+1 bivariate points (x0, y0), ... (xn, yn) in R²
// exists a unique polynomial of degree n passing through these points.
// Constructing these polynomials can be done in terms of Lagrange polynomials
func LangrangeInterpolation(y []int) func(int) float64{
	return func(z int) (result float64) {
		n := len(y)
		for i := 0; i < n; i++ {
			xi := i+1
			term := float64(y[i])
			for j := 0; j < n; j++ {
				xj := j+1
				if i != j {
					term *= float64(z - xj) / float64((xi - xj))
				}				
			}
			result += term
		}
		return result
	}
	
}

func SolveBoth(lines []string) (float64, float64, error) {
	var result1 float64
	var result2 float64
	for _, line := range lines {
		var values []int
		for _, symbol := range strings.Split(line, " ") {
			value, err := strconv.Atoi(symbol)
			if err != nil {
				return 0, 0, errors.New(fmt.Sprintf("Error parsing %s", symbol))
			}
			values = append(values, value)

		}
		polynomial := LangrangeInterpolation(values)
		n := len(values)
		result1 += polynomial(n+1)
		result2 += polynomial(0)
	}
	return result1, result2, nil	
}


func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	solutionPart1, solutionPart2, _ := SolveBoth(lines)
	fmt.Printf("Solution Part 1: %.2f\n", solutionPart1)
	fmt.Printf("Solution Part 2: %.2f\n", solutionPart2)
}