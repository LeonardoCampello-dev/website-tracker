package main

import "fmt"

func main() {
	userName := "Leonardo"
	appVersion := 0.1

	fmt.Println("Hello", userName)
	fmt.Println("Current version of the program", appVersion)

	fmt.Println("==============================")

	fmt.Println("(1) - start monitoring")
	fmt.Println("(2) - display the logs")
	fmt.Println("(0) - exit program")

	var chosenOption int

	fmt.Scan(&chosenOption)

	switch chosenOption {
	case 0:
		fmt.Println("Leaving")
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Logs")
	default:
		fmt.Println("Invalid option")
	}
}
