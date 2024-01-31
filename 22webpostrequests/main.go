package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Requests")

	myurl := "http://localhost:8080/post"

	requestBody := strings.NewReader(`
	{
		"message" : "This is post request"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	if response.StatusCode == 200 {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
