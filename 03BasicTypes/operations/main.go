package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 2, 4
	var sol1 int = (a + b) * (a - b)

	a++
	b += 10

	var sol2 int = a ^ b

	var r = 3.5

	var sol3 = math.Pi * r * r

	fmt.Printf(" solution 1 : %d \n solution 2 : %d \n solution 3 : %f \n", sol1, sol2, sol3)
}
