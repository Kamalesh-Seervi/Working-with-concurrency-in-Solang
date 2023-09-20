package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pizzaMade, pizzaFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNum int
	message  string
	success  bool
}

func (p *Producer) CLose() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
	// Thus now this logic is applied for producer struct
}

func makePizza(pizzaNum int) *PizzaOrder {
	pizzaNum++
	if pizzaNum <= NumofPizza {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recevied order #%d\n", pizzaNum)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzaMade++
		}
		total++
		fmt.Printf("Making pizza #%d\n. it will take %d seconds....\n",pizzaNum,delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <=2{
			
		}

	} else {

	}

}

func pizzaRia(pizzaMaker *Producer) {
	var i = 0
	for {
		currentPizza := makePizza(i)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println("hello")
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzaRia(pizzaJob)

}
