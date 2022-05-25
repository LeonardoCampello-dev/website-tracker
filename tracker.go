package main

import (
	"fmt"
	"os"
)

func main() {
	userGreetingMessage()

	displayInitialMenu()

	chosenOption := handleChosenOption()

	switch chosenOption {
	case 0:
		os.Exit(0)
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Logs")
	default:
		fmt.Println("Invalid option")

		os.Exit(-1)
	}
}

func userGreetingMessage() {
	userName := "Leonardo"
	appVersion := 0.1

	fmt.Println("Hello", userName)
	fmt.Println("Current version of the program", appVersion)

	fmt.Println("==============================")
}

func handleChosenOption() int {
	var chosenOption int

	fmt.Scan(&chosenOption)

	return chosenOption
}

func displayInitialMenu() {
	fmt.Println("(1) - start monitoring")
	fmt.Println("(2) - display the logs")
	fmt.Println("(0) - exit program")
}
