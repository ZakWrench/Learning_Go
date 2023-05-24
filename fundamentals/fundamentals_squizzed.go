package main

import "fmt"

func main() {
	// Data Types
	var name string = "John"
	age := 30

	// Variables and Constants
	const pi = 3.14159
	var radius float64 = 5.0

	// Operators and Control Structures
	area := pi * (radius * radius)
	if age >= 18 {
		fmt.Println(name, "is an adult.")
	} else {
		fmt.Println(name, "is a minor.")
	}

	// Arrays/Slices/Maps
	numbers := []int{1, 2, 3, 4, 5}
	names := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}

	// Functions and Recursion
	sum_2 := calculateSum(numbers...)
	factorial := calculateFactorial(5)

	fmt.Println("Area of the circle:", area)
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Sum of numbers:", sum_2)
	fmt.Println("Factorial of 5:", factorial)
}

func calculateSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func calculateFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * calculateFactorial(n-1)
}
