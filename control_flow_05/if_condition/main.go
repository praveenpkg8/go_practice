package main

import (
	"fmt"
)

func main() {

	var x = 2
	if x%2 == 0 {
		fmt.Printf(" %d is event \n", x)
	}

	var y = 5
	if y > 0 && y > 4 {
		fmt.Printf("%d is greater than 0 and 4 \n", y)
	}

	if z := -5; z < 0 {
		fmt.Printf("%d is negative \n", z)
	}
}
