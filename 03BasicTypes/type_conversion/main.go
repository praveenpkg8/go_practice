package main

import (
	"fmt"
)

func main() {

	var a int64 = 54
	var b int = int(a)
	var c float64 = 6.4

	var result = float64(b) * c

	var ranInt int = 93
	var ranUnit uint = uint(ranInt)
	var ranFloat float64 = float64(ranInt)

	fmt.Println(result, ranUnit, ranFloat)
}
