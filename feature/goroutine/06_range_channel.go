package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		fmt.Println("produce ", digit)
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit
		fmt.Println("calculate ", digit)
	}
	squareop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	go calcSquares(number, sqrch)
	data := <- sqrch
	fmt.Println("Final output",data)
}