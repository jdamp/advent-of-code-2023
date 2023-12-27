package main

// Function to calculate GCD using Euclidean algorithm
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Function to calculate LCM of two numbers
func lcm(a, b int) int {
    return a / gcd(a, b) * b // Use parentheses to ensure correct order of operations
}

// Function to calculate LCM of a slice of integers
func lcmSlice(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    result := numbers[0]
    for _, num := range numbers[1:] {
        result = lcm(result, num)
    }
    return result
}



