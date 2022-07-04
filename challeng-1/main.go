package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"
	wg.Add(3)
	go func() {
		updateMessage("Hello, universe!")
		printMessage()
		defer wg.Done()
	}()

	go func() {
		updateMessage("Hello, cosmos!")
		printMessage()
		defer wg.Done()
	}()
	go func() {
		updateMessage("Hello, world!")
		printMessage()
		defer wg.Done()
	}()
	wg.Wait()
}
