package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch and Case Statements")
	rand.NewSource(time.Now().UnixMicro())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The dice value is : ")
	switch diceNumber {
	case 1:
		fmt.Println(diceNumber)
	case 2:
		fmt.Println(diceNumber)
	case 3:
		fmt.Println(diceNumber)
	case 4:
		fmt.Println(diceNumber)
	case 5:
		fmt.Println(diceNumber)
	case 6:
		fmt.Print(diceNumber)
		fmt.Println("Get an other chance to role")
	default:
		fmt.Println("Whats that!")
	}

}
