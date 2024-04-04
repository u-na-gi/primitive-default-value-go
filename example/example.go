package main

import "fmt"

func main() {
	// This is a placeholder for the main package.

	var a string
	var b int
	var c float32
	var d bool
	var e complex64

	var f []int
	var g map[string]int

	fmt.Printf("string: %q\n", a == "")
	fmt.Printf("int: %d\n", b == 0)
	fmt.Printf("float32: %f\n", c == 0.0)
	fmt.Printf("bool: %t\n", d == false)
	fmt.Printf("complex64: %v\n", e == 0+0i)
	fmt.Printf("slice: %v\n", f == nil)
	fmt.Printf("map: %v\n", g == nil)

}
