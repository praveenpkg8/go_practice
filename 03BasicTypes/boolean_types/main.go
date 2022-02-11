package main

import "fmt"

func main() {
	var goBool bool = true
	var randomBool = false

	var checkBool = 5 >= 4
	var doubleCheckBool = 34 != 34

	var complexCheckBool = 5 == 5 && 4 < 4
	fmt.Println(goBool, randomBool, checkBool, doubleCheckBool, complexCheckBool)
}
