package main

import "fmt"

func main() {

	var number int

	number = 100

	if number > 0 {// this is an if statement. If the conditional evaluates to true, then we perform the code inside the body
		fmt.Println("the number is positive")
	} else if number < 0 {
		fmt.Println("the number is negative")
	} else {
		fmt.Println("the number is zero")
	}

	name := "Anthony"

	switch name { //this is a switch case. It check's for which case below equals the value in the variable then executes that code
	case "Anthony":// if the name variable equals "Anthony", then it will print his last name
		fmt.Println("Jordan-Parnell")
	case "Nick":
		fmt.Println("Winfield")
	default: // The default case runs if no other case is matched
		fmt.Println("Who?")
	}

}