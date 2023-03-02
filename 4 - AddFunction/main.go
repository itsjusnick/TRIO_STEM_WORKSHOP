package main

import "fmt"

func main() {
	var total int //this is how we declare new variables

	total = add(5, 10) //this is how we use a function

	fmt.Println(total)
}

func add(number1 int, number2 int) int { //this is how we declare new functions

	total := number1 + number2 //adding two integers (numbers)

	return total //we use 'return' to exit a function and give back results
}
