package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	endpoints := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.mongodb.com",
		"https://www.stackoverflow.com",
		"https://www.irctc.co.in",
		"https://jeemain.ntaonline.in",
	}
	a := time.Now().UnixMilli()
	wg.Add(len(endpoints) + 1)
	for _, web := range endpoints {
		go getStatusCode(web)
	}
	defer getTime(&a)
	wg.Wait()
}

func getTime(a *int64) {
	b := time.Now().UnixMilli()
	fmt.Println(b - *a)
}

func getStatusCode(endpoint string) {
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(endpoint, response.StatusCode)
	}
	defer wg.Done()
}
