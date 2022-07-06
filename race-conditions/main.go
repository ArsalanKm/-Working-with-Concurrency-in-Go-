package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	fmt.Println("In Function!")
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "hello world !"
	var mutex sync.Mutex
	wg.Add(2)
	go updateMessage("hello world! 2", &mutex)
	go updateMessage("hello world! 3", &mutex)
	wg.Wait()
	fmt.Println(msg)
}
