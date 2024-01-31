package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Lets start files in golang")

	filename, err := os.Create("./dheeraj.txt")
	defer filename.Close()
	if err != nil {
		fmt.Println(err)
	}
	length, err := filename.WriteString("Hi I am Dheeraj Sai Gogineni")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(length)
	cont, err := os.ReadFile("./dheeraj.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(cont))
	}
}
