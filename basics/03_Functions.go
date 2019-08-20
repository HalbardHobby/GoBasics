package main

import "fmt"

// A function can take any number of parameters
// Notice that type comes after variable name
// When 2 or more consecutive parameters shar a type,
// you can omit the type from all baut last.
// Hence (x int , y int) becomes (x, y int)
func add(x, y int) int {
	return x + y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Return values can be named. If so,
// they are treated as variables defined at the top of the function
// These names should bes used to document the meaning of return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 20))

	a, b := swap("Hello", "世界")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
