package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {
	for {
		s, ok := <-ping
		if !ok {
			// do some thing
			// wether the value sent to channel is zero value or not
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string)
	pong := make(chan string)
	//
	go shout(ping, pong)

	fmt.Println("Type something and press enter(enter q to quite)")

	for {
		fmt.Print("->")

		// get user input

		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Println("Response: ", response)

	}

	fmt.Println("All done closing channels")
	close(ping)
	close(pong)

}
