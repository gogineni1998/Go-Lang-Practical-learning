package main

import "fmt"

func main() {
	fmt.Println("Loops Work space")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i, day := range days {
		fmt.Println(i, day)
	}

	for i, day := range days {
		if i == 6 {
			goto printer
		}
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i, day)
	}

printer:
	fmt.Println("Tis is dheeraj")
}
