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
		color.Yellow("%s goes to the waiting room for check the clients", barber)
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
