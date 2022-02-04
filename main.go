package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter a number greater than 1 to play the game!")

	if len(os.Args) != 2 {
		fmt.Println("Please enter a number")
		return
	}

	userInput := os.Args[1]

	if userInput, err := strconv.Atoi(userInput); err != nil {
		fmt.Printf("%v is not a number\n", userInput)
	} else if userInput%2 == 0 && userInput%8 == 0 {
		fmt.Printf("%v is an even number and is divisible by 8\n", userInput)
	} else if userInput%2 == 0 {
		fmt.Printf("%v is an even number\n", userInput)
	} else if userInput%2 != 0 {
		fmt.Printf("%v is an odd number\n", userInput)
	}
}
