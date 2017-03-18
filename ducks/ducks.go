package main

import (
	"fmt"
)

type FlyBehavior interface {
	Fly()
}

type QuackBehavior func()


type IDuck interface {
	Quack()
	Fly()
}

type Duck struct {
	flyB FlyBehavior
	quackB QuackBehavior
}

type MallardDuck struct {
	Duck
}

type RubberDuck struct {
	Duck
}

func (d Duck) Quack() {
	d.quackB()
}

type flyNowWay struct {}
func (flyNowWay) Fly() {
	fmt.Println("I can't fly :(")
}
type flyWithWings struct {}
func (flyWithWings) Fly() {
	fmt.Println("Yay! I fly!")
}

func Squeeze() {
	fmt.Println("I squeeze")
}

func QuackWithMouth() {
	fmt.Println("I quack! I quack!")
}

func (d Duck) Fly() {
	d.flyB.Fly()
}

func playWithDuck(d IDuck) {
	fmt.Println(d)
	d.Fly()
	d.Quack()
}

func NewRubberDuck() *RubberDuck {
	d := &RubberDuck{}
	d.flyB = flyNowWay{}
	d.quackB = Squeeze
	return d
}

func NewMallardDuck() *MallardDuck {
	d := &MallardDuck{}
	d.flyB = flyWithWings{}
	d.quackB = QuackWithMouth
	return d
}

func main() {
	firstDuck := NewRubberDuck()
	playWithDuck(firstDuck)
	secondDuck := NewMallardDuck()
	playWithDuck(secondDuck)
}