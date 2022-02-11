package main

import "fmt"

func main() {
	var greeting string = "hello world"
	var luckyNumber int = 8
	var randomDecimal float64 = 5.6
	fmt.Println(greeting, luckyNumber, randomDecimal)

	//====================================

	var (
		rollNumber int    = 54
		comicHero  string = "BatMan"
	)
	fmt.Println(rollNumber, comicHero)

	//====================================

	series := "friends"
	rank, dailouge := 2, "I be there for you"
	fmt.Println(series, rank, dailouge)
}
