package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myUrl string = "http://localhost:8080/get"

func main() {
	fmt.Println("Perform get requests on spring rest api")
	getWebRequest()
}

func getWebRequest() {
	response, err := http.Get(myUrl)
	checkError(err)
	defer response.Body.Close()

	var buffer strings.Builder

	data, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println("The response of the API is : ", string(data))
	buffer.Write(data)
	fmt.Println(buffer.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
