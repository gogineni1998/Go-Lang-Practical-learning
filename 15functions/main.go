package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to functions world")
	getInfo()
	result := addNumbers(2, 3)
	fmt.Println(result)
	result, information := addNumbersWithInfo(2, 3, 4, 5, 6, 7, 8)
	fmt.Println(result, information)
}

func addNumbersWithInfo(numbers ...int) (int, string) {
	res := 0
	for _, number := range numbers {
		res = res + number
	}
	return res, "This is the result obtained"
}

func addNumbers(i1 int, i2 int) int {
	return i1 + i2
}

func getInfo() {
	fmt.Println("I am the first function declared here")
}
