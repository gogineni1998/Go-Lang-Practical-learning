package main

import "fmt"

func main() {
	var username string = "Dheeraj"
	fmt.Println(username)
	fmt.Printf("The type of this variable is %T \n", username)

	var username1 = "Dheeraj"
	fmt.Println(username1)
	fmt.Printf("The type of this variable is %T \n", username1)

	username2 := "Dheeraj"
	fmt.Println(username2)
	fmt.Printf("The type of this variable is %T \n", username2)

	var id int8 = 31
	fmt.Println(id)

}
