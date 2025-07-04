package main

import (
	"fmt"
	"local/mymath"
)

func main() {
	a, b := 5, 3

	sum := mymath.Add(a, b)
	product := mymath.Multiply(a, b)

	fmt.Printf("Add(%d, %d) = %d\n", a, b, sum)
	fmt.Printf("Multiply(%d, %d) = %d\n", a, b, product)
}
