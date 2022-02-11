package main

import "fmt"

func main() {
	var (
		name   string  = "El 'professor'"
		age    int     = 54
		salary float64 = 54000.546
	)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("salary: %.2f", salary)
}
