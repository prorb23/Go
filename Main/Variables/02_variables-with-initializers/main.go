package main

import (
	"fmt"
)

var i,j int = 1, 2

func main() {
	var c , java , python = true , true ,"!no"
	fmt.Println(i,j,c,java,python)
}

/*
	A var declaration can include initializers, one per variable.
	If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/