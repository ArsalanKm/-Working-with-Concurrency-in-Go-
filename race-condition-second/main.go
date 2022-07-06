package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var mutex sync.Mutex

	// print starting values
	fmt.Printf("initial account balance:$%d.00", bankBalance)
	fmt.Println()
	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Second Job", Amount: 10},
		{Source: "Third Job", Amount: 50},
		{Source: "Fourth Job", Amount: 100},
	}
	wg.Add(len(incomes))
	// loop through 52 week and print out how much is made keep a running total
	for i, income := range incomes {
		go func(i int, incom Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mutex.Lock()
				bankBalance += incom.Amount
				mutex.Unlock()
				fmt.Printf("On week %d, you earned %d.00 from %s\n", week, incom.Amount, incom.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance

	fmt.Printf("Final bank balance:%d", bankBalance)
}
