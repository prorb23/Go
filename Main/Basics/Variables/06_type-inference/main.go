package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
}

/*
	when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant
*/