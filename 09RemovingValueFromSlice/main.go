package main

import "fmt"

func main() {
	mySlice1 := []int{}
	mySlice1 = append(mySlice1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(mySlice1)
	ind := 5

	fmt.Println("Deleting value at index ", ind)
	fmt.Println(append(mySlice1[0:ind], mySlice1[ind+1:]...))
}
