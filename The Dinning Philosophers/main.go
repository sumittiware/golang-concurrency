package main

import (
	"fmt"
	"sync"
	"time"
)


type Philosopher struct {
	name string
	rightFork int
	leftFork int
}

var philosophers = []*Philosopher{
	{name: "Pluto", rightFork: 4, leftFork: 0},
	{name: "Socrates", rightFork:0, leftFork: 1},
	{name: "Aristotle", rightFork: 1, leftFork: 2},
	{name: "Pascal", rightFork: 2, leftFork: 3},
	{name: "Locke", rightFork: 3, leftFork: 4},
}

const hunger = 3

var eatTime = 0 *  time.Second
var thinkTime = 0 * time.Second
var sleepTime = 0 * time.Second

// This will be used to hold the order of the philosophers finishing their meal
var orderMutex = &sync.Mutex{}
var orderFinished []string

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("Philosophers are hungry and thinking")

	dine();

	fmt.Println("Philosophers are full and left the table!")
	fmt.Println("----------------------------")
	fmt.Println("Order of philosophers finishing their meal : ")

	for i, name := range orderFinished {
		fmt.Printf("%d. %s\n", i+1, name)
	}
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go dinningProblem(philosophers[i], forks, wg, seated)
	}

	wg.Wait()
}

func dinningProblem(p *Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", p.name)
	seated.Done()

	seated.Wait()

	for i := 0; i < hunger; i++ {
		fmt.Printf("%s is hungry\n", p.name)
		pickUpForks(p, forks)
		fmt.Printf("%s is eating\n", p.name)
		time.Sleep(eatTime)
		fmt.Printf("%s is thinking\n", p.name)
		time.Sleep(thinkTime)
		putDownForks(p, forks)
	}

	time.Sleep(sleepTime);
	fmt.Printf("%s is going to sleep\n", p.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, p.name)
	orderMutex.Unlock()
}

func pickUpForks(p *Philosopher, forks map[int]*sync.Mutex) {
	if(p.leftFork > p.rightFork) {
		forks[p.rightFork].Lock()
		forks[p.leftFork].Lock()
	}else {
		forks[p.leftFork].Lock()
		forks[p.rightFork].Lock()
	}
}

func putDownForks(p *Philosopher, forks map[int]*sync.Mutex) {
	forks[p.leftFork].Unlock()
	forks[p.rightFork].Unlock()
}