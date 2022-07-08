package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(time.Second * 6)
		ch <- "This is from server one"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server two"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("------------------")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("CASE ONE: ", s1)
		case s2 := <-channel1:
			fmt.Println("CASE TWO: ", s2)
		case s3 := <-channel2:
			fmt.Println("CASE Three: ", s3)
		case s4 := <-channel2:
			fmt.Println("CASE Four: ", s4)
		default:
			// avoiding deadlock when none of this channels are listening
			
		}
	}
}
