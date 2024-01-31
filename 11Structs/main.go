package main

import "fmt"

func main() {
	dheeraj := User{"Dheeraj", "gogineni1998@gmail.com", true, 16}
	fmt.Println(dheeraj)
	fmt.Printf("Full Details of Struct is : %+v\n", dheeraj)
	fmt.Printf("The name is %v and age is %+v", dheeraj.Name, dheeraj.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
