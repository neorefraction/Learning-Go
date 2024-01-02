package main

import "fmt"

func main() {
	// If else statement
	var name string

	fmt.Print("Introduce your name: ")
	fmt.Scan(&name)

	if name == "Johnny" {
    	fmt.Println("Welcome teacher!")
	} else if name == "Carlos" {
    	fmt.Println("Welcome dear Student!")
	} else {
    	fmt.Println("Unkown User. Access denied.")
	}

	// Using logical and relational operators
	var i, j int

	// Asks for two numbers
	fmt.Print("Introduce a number: ")
	fmt.Scan(&i)

	fmt.Print("Introduce another number: ")
	fmt.Scan(&j)

	// Evaluates the numbers
	if i > 0 && j > 0 {
	    fmt.Println("Both values are bigger than 0")
	}

	if i < 0 && j < 0 {
    	fmt.Println("Both values are lower than 0")
	}

	if i == 0 && j == 0 {
    	fmt.Println("Both values are lower than 0")
	}

	if i < 0 || j < 0 {
    	fmt.Println("One of the values is lower than 0")
	}
}