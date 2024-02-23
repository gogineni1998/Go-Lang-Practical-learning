package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now()
	ctx := context.Background()
	userId := 60
	fetchData(ctx, userId)
	fmt.Println(time.Since(t))
}

func fetchData(ctx context.Context, id int) {
	context.WithTimeout()
	val, err := fetchThirdPartyStuf()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", val)
}

func fetchThirdPartyStuf() (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 665, nil
}
