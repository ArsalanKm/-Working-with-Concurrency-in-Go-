package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel.")
		// simulate doing something
		time.Sleep(time.Second * 1)
	}
}
func main() {

	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("sending", i, "to channel")
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("Done !")

	close(ch)

}
