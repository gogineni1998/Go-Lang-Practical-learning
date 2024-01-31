package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func validate(rating string) (string, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(rating), 32)

	if err != nil {
		return "", err
	}
	if value < 0 || value > 5 {
		return "", errors.New("Please enter only between [0,5]")
	}

	return rating, nil
}

func formatTime(currentTime time.Time) string {
	return currentTime.Format("01-02-2006 Monday 15:04:05")
}

func main() {
	fmt.Println("Enter the rating stating from ")

	fmt.Println("Enter a input")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	rating, err := validate(input)

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Rating Validated : ", rating)
		fmt.Printf("Rating type is : %T\n", rating)
		currentTime := time.Now()
		formatedTIme := formatTime(currentTime)
		fmt.Println("Rating Stored in the meamory at time : ", formatedTIme)
		meamoryPtr := &formatedTIme
		fmt.Println(meamoryPtr)
		fmt.Println(*meamoryPtr)
	}
}
