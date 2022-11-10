package main

import "fmt"

func main() {

	total := add(5, 10)

	fmt.Println(total)
}

func add(number1 int, number2 int) int { //this is how we declare new functions
	var total int //this is how we declare new variables

	total = number1 + number2 //adding new integers (numbers)

	return total //we use 'return' to exit a function and give our results
}
