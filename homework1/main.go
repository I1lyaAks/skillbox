package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string
	for fmt.Scan(&str); str != "стоп"; fmt.Scan(&str) {
		number1 := make(chan int)
		number2 := make(chan int)
		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Что-то пошло не так, попробуйте снова")
			continue
		}
		go square(number1, number2)
		go multiplication(number2)
		number1 <- number
		close(number1)

	}
}
func square(number1, number2 chan int) {
	x := <-number1
	x *= x
	fmt.Println("квадрат:", x)
	number2 <- x
	close(number2)
}
func multiplication(number2 chan int) {
	x := <-number2
	x *= 2
	fmt.Println("умножение на 2:", x)
}
