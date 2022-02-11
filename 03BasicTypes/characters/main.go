package main

import "fmt"

func main() {
	var randomByte byte = 'a'
	var myEmoj rune = '♥'
	fmt.Printf("%c = %d and %c = %U \n", randomByte, randomByte, myEmoj, myEmoj)
}
