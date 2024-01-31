package main

import "fmt"

func main() {
	fmt.Println("Array Declaration Types : ")

	fmt.Println("Type 1 declaration : var myArray [4]string ")
	var myArray1 [4]string
	myArray1[0] = "Dheeraj"
	fmt.Println(myArray1)

	fmt.Println("Type 2 declaration : var myArray2 = [2]string{}")
	var myArray2 = [2]string{}
	myArray2[0] = "Dheeraj"
	fmt.Println(myArray2)

	fmt.Println("Type 3 declaration : myArray3 := [2]string{}")
	myArray3 := [2]string{}
	myArray3[0] = "Dheeraj"
	fmt.Println(myArray3)

	fmt.Println("Type 3 declaration : myArray4 := [...]string{}")
	myArray4 := [...]string{"Dheeraj"}
	fmt.Println(myArray4)

	fmt.Println("2d Arrays Declaration : myArray5 := [3][3]string{}")
	myArray5 := [3][3]string{}

	myArray5[0][0] = "Sai"
	myArray5[0][1] = "Sai"
	myArray5[0][2] = "Sai"
	myArray5[1][0] = "Sai"
	myArray5[1][1] = "Sai"
	myArray5[1][2] = "Sai"
	myArray5[2][0] = "Sai"
	myArray5[2][1] = "Sai"
	myArray5[2][2] = "Sai"
	fmt.Println(myArray5)

}
