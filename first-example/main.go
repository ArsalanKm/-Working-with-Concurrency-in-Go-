package main

import (
	"fmt"
	"sync"
)

func printSomething(s string,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	var wg sync.WaitGroup

	words := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eight",
	}

	wg.Add(len(words))

	for _, v := range words {
		go printSomething(v , &wg)
	}
	wg.Wait()

}
