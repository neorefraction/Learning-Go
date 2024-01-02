package main

import "fmt"

func main() {

	// It's possible to declare constants using 'const' and they must be initialized
	const pi float32 = 3.14
	const title string = "Constants in Go"

	fmt.Printf("Title: %s\n", title)
	fmt.Printf("Pi: %g\n", pi)
}
