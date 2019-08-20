package main

import "fmt"

// the 'var' statement declares a list of variables,
// as previously seen, the type is last
var c, python, java bool

// A 'var' declaration can include initializers, one for variable.
// If an initializer is present, the type can be omitted.
var j, k int = 1, 2

// A 'var' statement can be at package or function level.
func main() {
	var i int
	fmt.Println(i, j, k, c, python, java)

	// Inside a function the ':=' short assignment can be used
	// in place of a 'var' declaration with implicit type.
	l := "¬_¬"
	// Outside a function every statement begins with a keyword
	fmt.Println(l)
}
