package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	myurl := "http://localhost:8080/formdata"

	data := url.Values{}
	data.Add("message", "Hello")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	response_data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(response_data))

}
