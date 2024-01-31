package main

import (
	"fmt"
	"time"
)

func main() {

	presentTIme := time.Now()

	fmt.Println(presentTIme)
	formatedTime := presentTIme.Format("01-02-2006 15:04:05")
	fmt.Println(formatedTime)

	createdDate := time.Date(2021, 8, 15, 15, 3, 4, 0, time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
