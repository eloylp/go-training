package main

import (
	"fmt"
)

func main() {

	c1 := 12 + 1i
	c2 := complex(5, 7)

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)
	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

	x := 12
	k := 5

	fmt.Println(x)
	fmt.Printf("Type of x: %T\n", x)

	div := x / k
	fmt.Println("div", div)
}
