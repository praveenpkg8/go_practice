package main

import (
	"fmt"
)

func main() {
	var x = 3 + 5i
	var y = 5 + 6i
	fmt.Println(x, y)

	var q1 = 3.5
	var q2 = 4.6

	var q3 = complex(q1, q2)
	fmt.Println(q3)

	var sol1 = x + y
	var sol2 = x - y
	var sol3 = x * y
	var sol4 = x / y
	fmt.Println(sol1, sol2, sol3, sol4)

}
