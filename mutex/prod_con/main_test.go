package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestPizzaMaking(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	go pizzaRia(pizzaJob)

	numExpectedPizzas := NumofPizza
	numExpectedDeliveries := 0
	numExpectedFailures := 0

	for i := range pizzaJob.data {
		if i.pizzaNum <= NumofPizza {
			if i.success {
				numExpectedDeliveries++
			} else {
				numExpectedFailures++
			}
		} else {
			t.Fatal("Unexpected end of pizza making")
		}
	}

	if numExpectedDeliveries != pizzaMade {
		t.Fatalf("Expected %d successful deliveries, but got %d", numExpectedDeliveries, pizzaMade)
	}

	if numExpectedFailures != pizzaFailed {
		t.Fatalf("Expected %d failures, but got %d", numExpectedFailures, pizzaFailed)
	}

	if numExpectedPizzas != total {
		t.Fatalf("Expected %d total attempts, but got %d", numExpectedPizzas, total)
	}
}


func BenchmarkPizzaMaking(b *testing.B) {
    rand.Seed(time.Now().UnixNano())
    pizzaJob := &Producer{
        data: make(chan PizzaOrder),
        quit: make(chan chan error),
    }
    go pizzaRia(pizzaJob)

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        // Simulate pizza making and delivery
        <-pizzaJob.data
    }

    b.StopTimer()
    pizzaJob.CLose() // Close the pizza maker after the benchmark
}