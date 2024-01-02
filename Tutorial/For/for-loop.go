package main

import "fmt"

func main() {

	/* Classic for loops */

	// Classic for

	// It's necessary to declarate the loop counter
	var i int

	fmt.Print("For like loop: [")

	for i = 0; i < 4; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("]")

	// Go advanced loop
	fmt.Print("Advanced for like loop: [")

	for i := 0; i < 10; i++ {  // ':=' Assigns and infers the data type without specifying

		if i % 2 != 0 {
			continue  // It is possible to avoid iterations with 'continue'
		
		}
		fmt.Printf("%d ", i)
	}

	fmt.Println("]")

	/* While like loops */

	// While true
	for {
		fmt.Println("While True: []")
		break  // Necessary to avoid infinite loop
	}

	// Classic while
	var j int = 0

	fmt.Print("While like loop: [")
	
	for j < 10 {
		j += 1

		fmt.Printf("%d ", j)
	}

	fmt.Println("]")
}