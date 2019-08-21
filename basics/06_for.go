package main

import "fmt"

func main() {
	// Go only possesses for as a looping construct.
	// As usual it has a init statement,
	// a condition expression & a post statement
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Init and post statements are optional
	// Likewise For is Go's while
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
