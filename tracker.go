package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

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

	sites := readSitesFromFile()

	for index := 0; index < monitoring; index++ {
		fmt.Println("Test -", index+1)

		for _, site := range sites {
			testsWebsiteAvailability(site)
		}

		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
}

func testsWebsiteAvailability(site string) {
	response, error := http.Get(site)

	if error != nil {
		fmt.Println("some error occurred!", error)
	}

	if response.StatusCode == 200 {
		fmt.Println("Website", site, "was successfully loaded")
	} else {
		fmt.Println("Website", site, "is having problems. Status Code:", response.StatusCode)
	}
}

func readSitesFromFile() []string {
	var sites []string

	file, error := os.Open("sites.txt")

	if error != nil {
		fmt.Println("some error occurred!", error)
	}

	reader := bufio.NewReader(file)

	for {
		row, error := reader.ReadString('\n')

		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if error == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
