package main

import (
	"fmt"
	"strconv"
)

var names = []string{"Nick", "Anthony", "Marissa", "Lydia", "Lucas"} //Array - a list of values

func main() {
	
	sum := 0
	for  i := 0; i < 5; i++ { //for loop - will continue to do whatever is in the brackets until the conditions are met
		//fmt.Println(i)
		sum = sum + i
	}
	fmt.Println(sum)

	for index, name := range names { // for range - will loop through the elements of the array
		fmt.Println(name + " is in position " + strconv.Itoa(index) + " in line") //strconv.Itoa - converts the integer to a string
	}
}

