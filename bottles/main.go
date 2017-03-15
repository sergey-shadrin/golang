package main

import "fmt"

func getBottleString(bottleNumber int) string {
	bottleStringEnding := ""
	digit := bottleNumber % 10
	twoDigits := bottleNumber % 100
	if (twoDigits >= 11 && twoDigits <= 14) || ((digit == 0) || (digit >= 5 && digit <= 9)) {
		bottleStringEnding = "бутылок"
	} else if digit == 1 {
		bottleStringEnding = "бутылка"
	} else if digit >= 2 && digit <= 4 {
		bottleStringEnding = "бутылки"
	}
	return fmt.Sprintf("%v %v", bottleNumber, bottleStringEnding)
}

func main() {
	initialBottlesCount := 199
	result := ""
	for i := initialBottlesCount; i > 0; i-- {
		result += fmt.Sprintf("%v пива на стене\n%v пива!\n", getBottleString(i), getBottleString(i))
		result += fmt.Sprintf("Возьми одну, пусти по кругу\n")
		if (i - 1) != 0 {
			result += fmt.Sprintf("%v пива на стене\n\n", getBottleString(i-1))
		} else {
			result += "Нет бутылок пива на стене!\nНет бутылок пива!\nПойди в магазин и купи еще,\n"
			result += fmt.Sprintf("%v пива на стене!", getBottleString(initialBottlesCount))
		}
	}
	fmt.Println(result)
}
