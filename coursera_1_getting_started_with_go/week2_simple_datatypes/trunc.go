package main

import (
	"fmt"
	"strconv"
)

/*
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
*/

func main() {
	var userInput float64

	fmt.Printf("Enter a floating point number \n")

	_, err := fmt.Scan(&userInput)

	if err != nil {
		panic(err)
	}

	var userInputTrunc int = int(userInput)

	fmt.Printf(strconv.Itoa(userInputTrunc) + "\n")
}
