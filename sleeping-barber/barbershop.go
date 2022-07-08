package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow(" %s goes to the waiting room for check the clients", barber)
		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do,%s is going to take a nap", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}

				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s'hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	// The barber is going home and receive no more clients
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("=============================================================")
	color.Green("Ther barber shop is now closed for the day and every one has gone home.")
}

func (shop *BarberShop) addClient(client string) {
	// print out the message
	color.Green("*** %s arrives !", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in waiting room", client)
		default:
			color.Red("The waiting room is full so %s leaves", client)
		}

	} else {
		color.Red("The shop is already closed so %s leaves", client)
	}
}
