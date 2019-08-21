package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {
	// Go's if statements are very similar to C
	// just drop the ()
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Similar to 'for', the if statement can start with a
	// shorthand statement. Variables declared there will
	// only be scoped to the if body
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func os() string {
	// a Switch statement is basically a shorter way to write
	// if/else statements. It runs the first case whose value
	// is equal to the condition expression.
	switch os := runtime.GOOS; os {
	case "darwin":
		return "OS X."
	case "linux":
		return "Linux."
	default:
		// freebsd, openbsd,
		// plan9, windows...
		return os
	}
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("Go runs on ", os())
}
