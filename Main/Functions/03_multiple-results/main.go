package main

import (
	"fmt"
)

func swap(a , b string) (string,string){
	return b,a
}

func main() {
	a,b := swap("hello","world")
	fmt.Println(a,b)
}

/*
	A function can return any number of results.
	The swap function returns two strings.

	The line
		a, b := swap("hello", "world")

		What it does

		Calls the function swap("hello", "world")

		That function returns two values

		Those two returned values are assigned to a and b

		a and b are created (declared) right here
*/