package main

import "fmt"

func main() {

	// Basic array declaration.
	var my_array [3]string

	// Go sets the array with zero-valued values by default.
	fmt.Printf("Default array values: %v\n\n", my_array)

	// It is possible to set the arrays values using indexes. Indexes starts on 0.
	my_array[0] = "Apple"
	my_array[1] = "Orange"
	my_array[2] = "Banana"

	fmt.Printf("New array values: %v\n\n", my_array)

	// In-line array declaration
	var inline_array = [5]int{1, 5, 7, 9, -1}

	fmt.Printf("Inline array: %v\n\n", inline_array)

	// It's possible to create 2D arrays also
	var my_2d_array [2][2]string

	my_2d_array[0][0] = "Ford"
	my_2d_array[0][1] = "Mazda"
	my_2d_array[1][0] = "Audi"
	my_2d_array[1][1] = "Volkswagen"

	fmt.Printf("2D array: %v\n\n", my_2d_array)

	// It's possible to create it in-line too
	var inline_2d_array = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("Inline 2D array: %v\n\n", inline_2d_array)
}
