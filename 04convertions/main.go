package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	rating, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rating + 1)
	}

}
