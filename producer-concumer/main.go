package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error // one channel that contains channels of error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

type Consumer struct {
}

func (p *Producer) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch

}

func makePizza(i int) *PizzaOrder {
	i++
	if i <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number %d\n", i)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza %d It will take %d seconds \n", i, delay)
		time.Sleep(time.Duration(delay) * (time.Second))

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza %d", i)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making the pizza %d", i)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order %d is ready", i)
		}
		order := PizzaOrder{
			pizzaNumber: i,
			message:     msg,
			success:     success,
		}

		return &order
	}

	return &PizzaOrder{
		pizzaNumber: i,
	}

}

func pizzeria(pizzaMaker *Producer) {
	//  keep track of which pizza we are making
	i := 0

	// run for ever or until we receive a quit notification

	// try to make pizzass
	// this loop run for ever and send created pizza to the data channel
	// stop when quit channel has something
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make pizza (we sent something to data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

		// decision

	}

}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out message

	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order number %d is out for delivery", i)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad !")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	color.Cyan("----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but faile to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day")
	case pizzasFailed >= 4:
		color.Red("It was an okay")
	case pizzasFailed >= 2:
		color.Yellow("It was pretty good day")
	default:
		color.Yellow("It was pretty good day")

	}

}
