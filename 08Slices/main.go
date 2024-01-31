package main

import "fmt"

func main() {

	fmt.Println("Slices Declaration Type : 1 :- ")
	var mySlice1 = []int{}
	mySlice1 = append(mySlice1, 1, 2, 3, 4, 5)
	fmt.Println(mySlice1)

	fmt.Println("Slices Declaration Type : 2 :- ")
	mySlice2 := []int{}
	mySlice2 = append(mySlice2, 2, 3, 4, 5)
	fmt.Println(mySlice2)

	fmt.Println("Slicing the Slices : ")
	fmt.Println(mySlice1[0:2])

	fmt.Println("using make() operator")
	mySlice3 := make([]int, 0)
	mySlice3 = append(mySlice3, 1, 2, 3, 4, 5, 6)
	fmt.Println(mySlice3)
}
