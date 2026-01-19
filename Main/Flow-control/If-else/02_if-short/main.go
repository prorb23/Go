package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64{
	if v:= math.Pow(x,n); v < limit{
		return v;
	}
	return limit // can not use v here as it is recognized inside only the scope of the if statement
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}

/*
	Like for, the if statement can start with a short statement to execute before the condition.
	Variables declared by the statement are only in scope until the end of the if.
*/ 