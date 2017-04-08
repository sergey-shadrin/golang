package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	millisecondsToTakeFork = 1000
	millisecondsToReleaseFork = 1000
	maxMillisecondsToContemplate = 1000
	maxMillisecondsToEat = 1000
	forkValue = 1
)

type Fork chan int

type Philosopher struct {
	name string
	leftFork Fork
	rightFork Fork
}

func (phil *Philosopher) Contemplate() {
	fmt.Printf("%v starts contemplating\n", phil.name)
	timeToContemplate := rand.Intn(maxMillisecondsToContemplate)
	time.Sleep(time.Duration(timeToContemplate) * time.Millisecond)
	fmt.Printf("%v has contemplated for %v milliseconds\n", phil.name, timeToContemplate)
}

func (phil *Philosopher) Eat() {
	fmt.Printf("%v wants to eat\n", phil.name)
	phil.takeFork(phil.leftFork, "left")
	phil.takeFork(phil.rightFork, "right")
	phil.doEat()
	phil.releaseFork(phil.leftFork, "left")
	phil.releaseFork(phil.rightFork, "right")
}

func (phil *Philosopher) doEat() {
	fmt.Printf("%v eats spaghetti\n", phil.name)
	timeToEat := rand.Intn(maxMillisecondsToEat)
	time.Sleep(time.Duration(timeToEat) * time.Millisecond)
	fmt.Printf("%v has finished eating after %v milliseconds\n", phil.name, timeToEat)
}

func (phil *Philosopher) takeFork(fork Fork, forkName string) {
	fmt.Printf("%v is going to take %v fork\n", phil.name, forkName)
	time.Sleep(millisecondsToTakeFork * time.Millisecond)
	<-fork
	fmt.Printf("%v has taken %v fork\n", phil.name, forkName)
}

func (phil *Philosopher) releaseFork(fork Fork, forkName string) {
	fmt.Printf("%v is going to release %v fork\n", phil.name, forkName)
	time.Sleep(millisecondsToReleaseFork * time.Millisecond)
	fork <- forkValue
	fmt.Printf("%v has released %v fork\n", phil.name, forkName)
}

func (phil *Philosopher) HaveDinner() {
	for {
		phil.Contemplate()
		phil.Eat()
	}
}

func main() {
	forks := [5]Fork{}
	for i := 0; i < len(forks); i++ {
		forks[i] = make(Fork, 1)
		forks[i] <- forkValue
	}

	philosophers := [5]Philosopher{
		{"Aristotle", forks[0], forks[1]},
		{"Nietzsche", forks[1], forks[2]},
		{"Kant", forks[2], forks[3]},
		{"Freud", forks[3], forks[4]},
		{"Plato", forks[4], forks[0]},
	}
	for i := 0; i < len(philosophers); i++ {
		go philosophers[i].HaveDinner()
	}

	time.Sleep(time.Minute)
}