package main

import "fmt"

func main() {
	var first, last string
	fmt.Println("Type your first name")
	fmt.Scanln(&first) //fmt.Scanln takes in input from console
	fmt.Println("Type your last name")
	fmt.Scanln(&last)

	fmt.Println("Hello " + name(first, last) + ", nice to meet you!")

}

func name(first string, last string) string {
	var wholeName string

	wholeName = first + " " + last //This is called concatenation

	return wholeName
}
