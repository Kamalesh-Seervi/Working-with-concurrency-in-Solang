package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amt    int
}

func main() {
	var bankBalance int
	var balance sync.Mutex
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{
			Source: "Job",
			Amt:    25000,
		},
		{
			Source: "FreeLance",
			Amt:    10000,
		},
		{
			Source: "Giftcard",
			Amt:    1000,
		},
		{
			Source: "Momgifts",
			Amt:    2000,
		},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amt
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d,you earned $%d.00 from %s\n", week, income.Amt, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	fmt.Printf("Final bank balance:$%d.00", bankBalance)
	fmt.Println()
}
