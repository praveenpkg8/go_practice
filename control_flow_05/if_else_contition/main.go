package main

import (
	"fmt"
)

func main() {

	var x int
	fmt.Println("enter a number")
	fmt.Scanf("%d", &x)
	if x%2 == 0 {
		fmt.Printf("%d is even \n", x)
	} else {
		fmt.Printf("%d is odd \n", x)
	}
}
