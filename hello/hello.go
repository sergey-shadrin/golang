package main

import (
	"fmt"
	"math"
)

const Epsilon = 0.00000000001

func swap(first, second int) (int, int) {
	return second, first
}

func countDown() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func whileInGo() (loopResult int) {
	for loopResult < 1000 {
		loopResult += loopResult + 1
	}
	return loopResult
}

func equals(first, second float64) {
	if delta := math.Abs(second - first); delta < Epsilon {
		fmt.Printf("Equals! Delta is %v\n", delta)
	} else {
		fmt.Printf("Not equals! Delta is %v\n", delta)
	}
}

func printTypes() {
	var (
		a int = 5
		b int8 = 2
		c int32 = 256
		d int64 = -145412341234
		e float32 = 132.1234123
		f float64 = 2423452.12341234
		g string = "Mamma mia!"
		h rune = 'Ю'
		i rune = 143123782
		j complex64 = 3 + 2i
		k complex128 = -5 + 29i
		l bool = false
		m bool = true
	)
	fmt.Println("Here we go with types:")
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Printf("%T: %v\n", c, c)
	fmt.Printf("%T: %v\n", d, d)
	fmt.Printf("%T: %v\n", e, e)
	fmt.Printf("%T: %v\n", f, f)
	fmt.Printf("%T: %v\n", g, g)
	fmt.Printf("%T: %v\n", h, h)
	fmt.Printf("%T: %v\n", i, i)
	fmt.Printf("%T: %v\n", j, j)
	fmt.Printf("%T: %v\n", k, k)
	fmt.Printf("%T: %v\n", l, l)
	fmt.Printf("%T: %v\n", m, m)
	fmt.Printf("%T: countDown\n", countDown)
}

func pointers() {
	fmt.Println("Pointers")
	i := 5
	j := 92233
	p := &i
	fmt.Println(*p)
	*p = j
	fmt.Println(i)

	p = &j
	j *= 6
	fmt.Println(*p)
}

func main() {
	fmt.Println("hello, world\n")
	fmt.Println(swap(5, 6))
	countDown()
	fmt.Println(whileInGo())
	equals(0.5, 0.7)
	equals(0.5, float64(1) / float64(2))
	printTypes()
	pointers()
}