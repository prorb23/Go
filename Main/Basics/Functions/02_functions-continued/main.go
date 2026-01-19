package main

import (
	"fmt"
)

func add(x , y , z int) int{
	return x+y+z
}

func main() {
	fmt.Println(add(10,10,10))
}

/*
	When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	In this example, we shortened
			x int, y int , z int
			to
			x, y, z int
*/