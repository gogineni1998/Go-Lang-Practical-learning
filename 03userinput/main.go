package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadSlice(1)

	fmt.Print(input, err)
	fmt.Printf("Type of the input is %T", input)
}
