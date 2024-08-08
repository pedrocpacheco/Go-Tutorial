package main

import (
	"fmt"
)

func main() {

	// Declare and initialize a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Print the original slice
	fmt.Println("Original slice:", numbers)

	// Append a new element to the slice
	numbers = append(numbers, 6)

	// Print the updated slice
	fmt.Println("Updated slice:", numbers)

	// Get the length of the slice
	length := len(numbers)

	// Print the length of the slice
	fmt.Println("Length of the slice:", length)

	// Get the capacity of the slice
	capacity := cap(numbers)

	// Print the capacity of the slice
	fmt.Println("Capacity of the slice:", capacity)

	// Get the value at index 2
	value := numbers[2]

	// Print the value at index 2
	fmt.Println("Value at index 2:", value)

	// Remove the element at index 2
	numbers = append(numbers[:1], numbers[3])
	fmt.Println(numbers)

}
