package main

import (
	"fmt"
	"sync"
	"time"
)

const hunger = 3

// variables - philosophrs
var philosophers = []string{"Arsalan", "Sahand", "Nazanin", "Shahriyar", "Negin"}
var wg sync.WaitGroup

var sleepTime = time.Second * 1
var eatTime = time.Second * 3
var thinkTime = time.Second * 1
var orderMutex sync.Mutex

var finished = make([]string, 3)

func diningProblem(philosopher string, leftHand, rightHand *sync.Mutex) {
	defer wg.Done()

	// print a message
	fmt.Printf("Philosopher %s is sat on the table\n", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		fmt.Println(philosopher, "isHungry")
		time.Sleep(sleepTime)

		// lock both forks
		leftHand.Lock()
		fmt.Printf("\t %s picked up the fork to his left\n", philosopher)
		rightHand.Lock()
		fmt.Printf("\t %s picked up the fork to his right\n", philosopher)

		// print a message
		fmt.Println(philosopher, "has both forks and he is eating")
		time.Sleep(eatTime)

		// time to think

		fmt.Println(philosopher, "is thinking")
		time.Sleep(thinkTime)
		// unlock the mutexes

		leftHand.Unlock()
		fmt.Printf("\t%s put down the fork on his left\n", philosopher)
		rightHand.Unlock()
		fmt.Printf("\t%s put down the fork on his right\n", philosopher)

		time.Sleep(sleepTime)
	}
	fmt.Println(philosopher, "is satisfied")
	time.Sleep(sleepTime)
	fmt.Println(philosopher, "has left the table.")
	orderMutex.Lock()
	finished = append(finished, philosopher)
	orderMutex.Unlock()

}

func main() {
	// print intro

	fmt.Println("The Dining Philosophers Problem")
	fmt.Println("--------------------------------")

	// spawn one goroutine for each philosophers
	wg.Add(len(philosophers))
	forkleft := &sync.Mutex{}
	for i := 0; i < len(philosophers); i++ {
		// create mutext for right fork
		forkRight := &sync.Mutex{}
		// call a goroutine
		go diningProblem(philosophers[i], forkleft, forkRight)
		forkleft = forkRight
	}
	wg.Wait()

	fmt.Println("The Table is Empty!")
	fmt.Print(finished)
}
