package main

import "fmt"

func swap(first, second int) (int, int) {
	return second, first
}

func countDown() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println("hello, world\n")
	fmt.Println(swap(5, 6))
	countDown()
}