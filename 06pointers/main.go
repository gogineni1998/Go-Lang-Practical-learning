package main

import "fmt"

func main() {
	fmt.Println("hi")

	myNum := 25
	ptr := &myNum
	ptr2 := &ptr
	ptr3 := &ptr2

	fmt.Println("value of ptr3 is address of ptr2", ptr3)

	fmt.Println("value at address in ptr3", *ptr3)

	fmt.Println("value at address of *ptr3", **ptr3)

	fmt.Println("value at address of **ptr3", *(*(*(ptr3))))
}
