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
	quit chan chan error 
}

type PizzaOrder struct {
	pizzaNumber int 
	message string
	success bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++;

	if pizzaNumber < NumberOfPizzas {
		delay := rand.Intn(5)+1
		fmt.Printf("Recieved Order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1

		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++;
		}else {
			pizzasMade++;
		}
		total++;

		fmt.Printf("Making Pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)	
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** Pizza #%d is burnt!", pizzaNumber);
		}else if rnd <= 4 {
			msg = fmt.Sprintf("*** Pizza #%d is undercooked!", pizzaNumber);
		}else {
			msg = fmt.Sprintf("*** Pizza #%d is ready!", pizzaNumber);
			success = true;
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzazone(pizzaMaker *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			select {
				// We tried to make the pizza
				case pizzaMaker.data <- *currentPizza:

				case quitChan := <-pizzaMaker.quit:
					close(pizzaMaker.data)
					close(quitChan)
					return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("The Pizzazone is opening for the business!")
	color.Cyan("------------------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzazone(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber < NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!\n", i.pizzaNumber)
			}else {
				color.Red(i.message)
				color.Red("Order #%d is cancelled!\n", i.pizzaNumber)
			}
		} else {
			color.Cyan("The Pizzazone is closing for the day!")
			err := pizzaJob.Close()

			if err != nil {
				color.Red("*** Error closing channel: %v", err)
			}
		}
	}

	color.Cyan("------------------------------------------")
	color.Cyan("Done for the day!")

	switch {
	case pizzasFailed > 9:
		color.Red("We had a awful day! Only %d pizzas made out of %d orders!", pizzasMade, total)
	case pizzasFailed >= 6:
		color.Red("We had a very bad day! %d pizzas made out of %d orders!", pizzasMade, total)
	case pizzasFailed >= 4:
		color.Yellow("We had a bad day! %d pizzas made out of %d orders!", pizzasMade, total)
	case pizzasFailed >= 2:
		color.Yellow("We had a good day! %d pizzas made out of %d orders!", pizzasMade, total)
	default:
		color.Green("We had a great day! %d pizzas made out of %d orders!", pizzasMade, total)
	}
	
}