package main

import "fmt"
import "anagram/anagram"

func main() {
	isAnagram := anagram.IsAnagram("Привет", "Тевирп")
	fmt.Println(isAnagram)
}