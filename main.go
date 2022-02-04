package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter a numbere to play the game!")

	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
		return
	}

	userInput := os.Args[1]

	if userInput, err := strconv.Atoi(userInput); err != nil {
		fmt.Printf("%v is not a number", userInput)
	} else if userInput%2 == 0 && userInput%8 == 0 {
		fmt.Printf("%v is an even number and is divisible by 8", userInput)
	} else if userInput%2 == 0 {
		fmt.Printf("%v is an even number", userInput)
	} else if userInput%2 != 0 {
		fmt.Printf("%v is an odd number", userInput)
	}
}
