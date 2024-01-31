package main

import "fmt"

func main() {
	fmt.Println("Executing Defer Keyword")
	defer fmt.Println("Finished Executing")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer seeDefer()

}

func seeDefer() {
	fmt.Println("Get defered")
}
