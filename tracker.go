package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	userGreetingMessage()

	for {
		displayInitialMenu()

		chosenOption := handleChosenOption()

		switch chosenOption {
		case 0:
			os.Exit(0)
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs")
		default:
			fmt.Println("Invalid option")

			os.Exit(-1)
		}
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

func startMonitoring() {
	fmt.Println("Monitoring")

	site := "https://random-status-code.herokuapp.com/"

	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Website", site, "was successfully loaded")
	} else {
		fmt.Println("Website", site, "is having problems. Status Code:", response.StatusCode)
	}
}
