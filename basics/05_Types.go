package main

import (
	"fmt"
)

// Go's Basic types are
var (
	b bool

	s string

	i8  int
	i16 int16
	i32 int32
	i64 int64

	ui   uint
	ui8  uint8
	ui16 uint16
	ui32 uint32
	ui64 uint64
	uip  uintptr

	by byte // alias for uint8

	rn rune // alias for int32
	// represents unicode code point

	fl32 float32
	fl64 float64

	cmp64  complex64
	cmp128 complex128
)

// The expression T(v) converts value v to type T
var i = 42
var f = float32(i)
var u = uint(f)

func main() {
	// Variables declared without an explicit initial value
	// are given their "zero" value.
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", s, s)
	fmt.Printf("Type: %T Value: %v\n", i8, i8)
	fmt.Printf("Type: %T Value: %v\n", i16, i16)
	fmt.Printf("Type: %T Value: %v\n", i32, i32)
	fmt.Printf("Type: %T Value: %v\n", ui64, ui64)
	fmt.Printf("Type: %T Value: %v\n", ui, ui)
	fmt.Printf("Type: %T Value: %v\n", ui8, ui8)
	fmt.Printf("Type: %T Value: %v\n", ui16, ui16)
	fmt.Printf("Type: %T Value: %v\n", ui64, ui64)
	fmt.Printf("Type: %T Value: %v\n", uip, uip)
	fmt.Printf("Type: %T Value: %v\n", by, by)
	fmt.Printf("Type: %T Value: %v\n", rn, rn)
	fmt.Printf("Type: %T Value: %v\n", fl32, fl32)
	fmt.Printf("Type: %T Value: %v\n", fl64, fl64)
	fmt.Printf("Type: %T Value: %v\n", cmp64, cmp64)
	fmt.Printf("Type: %T Value: %v\n", cmp128, cmp128)
}
