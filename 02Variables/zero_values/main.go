package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age                 int
		isConfirmed         bool
		salary              float64
	)

	fmt.Printf("%s %s %d %t %f", firstName, lastName, age, isConfirmed, salary)
}
