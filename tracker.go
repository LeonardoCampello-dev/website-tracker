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

	if chosenOption == 1 {
		fmt.Println("Monitoring")
	} else if chosenOption == 2 {
		fmt.Println("Logs")
	} else if chosenOption == 0 {
		fmt.Println("Leaving")
	} else {
		fmt.Println("invalid option")
	}
 }