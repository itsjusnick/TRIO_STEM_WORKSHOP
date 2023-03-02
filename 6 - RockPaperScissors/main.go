package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var playerChoice string
	var Choices = [3]string{"rock", "paper", "scissors"}
	again := true

	for again {

		fmt.Println("Please type your choice: rock, paper, scissors")
		fmt.Scanln(&playerChoice) //fmt.Scanln takes in input from console

		randomNumber := rand.Intn(3)
		botChoice := Choices[randomNumber] //this how we get a specfic value from an array

		fmt.Println("You chose " + playerChoice)
		fmt.Println("Bot chose " + botChoice)

		if botChoice == "paper" && playerChoice == "rock" {
			fmt.Println("You lose!")
		} else if botChoice == "paper" && playerChoice == "scissors" {
			fmt.Println("You win!")
		} else if botChoice == "rock" && playerChoice == "scissors" {
			fmt.Println("You lose!")
		} else if botChoice == "rock" && playerChoice == "paper" {
			fmt.Println("You win!")
		} else if botChoice == "scissors" && playerChoice == "rock" {
			fmt.Println("You win!")
		} else if botChoice == "scissors" && playerChoice == "paper" {
			fmt.Println("You lose!")
		} else {
			fmt.Println("Draw!")
		}

		fmt.Println("Would you like to play again? y/n")

		fmt.Scanln(&playerChoice)

		if playerChoice == "y" {
			again = true
		} else {
			again = false
		}
	}
}
