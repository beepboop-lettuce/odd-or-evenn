package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please enter a number greater than 1")
		return
	}

	userInt := os.Args[1]
	//	userInput is the result of userInt converted by Atoi into an int
	// userInt this is the user input
	if userInput, err := strconv.Atoi(userInt); err != nil {
		// we have userInt after the comma because we want the output message to refer to the user inputted value
		// before the value after the decimal has been dropped.
		fmt.Printf("%s is not a number\n", userInt)
	} else if userInput%2 == 0 && userInput%8 == 0 {
		fmt.Printf("%v is an even number and is divisible by 8\n", userInput)
	} else if userInput%2 == 0 {
		fmt.Printf("%v is an even number\n", userInput)
	} else if userInput%2 != 0 {
		fmt.Printf("%v is an odd number\n", userInput)
	}
}
