package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumofPizza = 10

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
		fmt.Printf("Making pizza #%d\n. it will take %d seconds....\n", pizzaNum, delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran of ingredients #%d:", pizzaNum)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza $%d!", pizzaNum)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza Order Success #%d!", pizzaNum)
		}

		p := PizzaOrder{
			pizzaNum: pizzaNum,
			message:  msg,
			success:  success,
		}
		return &p

	}
	return &PizzaOrder{
		pizzaNum: pizzaNum,
	}

}

func pizzaRia(pizzaMaker *Producer) {
	defer close(pizzaMaker.data)
	var i = 0
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNum
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.quit)
				close(quitChan)
				return
			}
		}
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzaRia(pizzaJob)

	// Consumer create and run it

	for i := range pizzaJob.data {
		if i.pizzaNum <= NumofPizza {
			if i.success {
				fmt.Printf("Order is out for delivery %d \n", i.pizzaNum)
			} else {
				fmt.Println("Customer is angry")
			}
		} else {
			fmt.Println("Done for the day")
			err := pizzaJob.CLose()
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("we made %d pizzas,but failed to make %d, with %d attempts in total.\n", pizzaMade, pizzaFailed, total)
}
