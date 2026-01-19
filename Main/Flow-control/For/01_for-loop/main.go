package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i:= 0; i < 10; i++ {
		sum += i;
	}
	fmt.Println(sum)
}

/*
	Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
*/