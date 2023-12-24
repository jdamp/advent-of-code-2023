package main

import (
	"fmt"
	"math"
)

// To solve a race with distance d and record t, we can vary the pressure time
// p to speed up the boat to a velocitity v=p.
// The total distance traveled d_t is then given as:
// d_t = v * t_t = v * (t - p) = p * (t - p)
// To win a race, d_t > d is required. This results in the quadratic inequality
// p**2 - p*t + d < 0, which has the zeros
// p_0,1 = t/2 +- sqrt((t/2)**2 - d)
//
// All input times p fulfilling p_0 < p < p1 will mean that we are winning the race.
// Therefore, solving part 1 means to find the zeros and check for all integer falling
// between the two zeros
func SolveRace(t, d float64) (p0, p1 float64) {
	eps := 1e-9
	p0 =t/2. - math.Sqrt(math.Pow(t/2., 2) - float64(d)) + eps
	p1 =t/2. + math.Sqrt(math.Pow(t/2., 2) - float64(d)) - eps
	return p0, p1
}

func SolvePart1(times []int, distances []int) int{
	marginOfError := 1
	for i := 0; i < len(times); i++ {
		p0, p1 := SolveRace(float64(times[i]), float64(distances[i]))
		nWaysToWin := int(math.Floor(p1) - math.Ceil(p0)) +1
		marginOfError *= int(nWaysToWin)
	}
	return marginOfError
}

func SolvePart2(time int, distance int) int{
	p0, p1 := SolveRace(float64(time), float64(distance))
	nWaysToWin := int(math.Floor(p1) - math.Ceil(p0)) +1
	return nWaysToWin
}

func main() {
	times := []int {41, 96, 88, 94}
	distances := []int {214, 1789, 1127, 1055}
	marginOfError := SolvePart1(times, distances)
	nWaysToWin := SolvePart2(41968894, 214178911271055)

	fmt.Printf("Part 1: %d\n", marginOfError)
	fmt.Printf("Part 2: %d\n", nWaysToWin)

}