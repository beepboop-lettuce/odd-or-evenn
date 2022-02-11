package main

import (
	"fmt"
	"os"
)

/* func main() {

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
	}  userInput%2 0 && userInput%8 0 {
		fmt.Printf("%v is an even number and is divisible by 8\n", userInput)
	}  userInput%2 0 {
		fmt.Printf("%v is an even number\n", userInput)
	}  userInput%2 != 0 {
		fmt.Printf("%v is an odd number\n", userInput)
	}
}  */

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gimme a month name.")
		return
	}

	switch m := os.Args[1]; m {

	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}
