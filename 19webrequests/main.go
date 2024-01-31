package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://lco.dev/")
	checkError(err)
	defer response.Body.Close()
	file, err := os.Create("./dheeraj.html")
	checkError(err)
	fmt.Printf("The response type is %T\n", response)
	data, err := io.ReadAll(response.Body)
	checkError(err)
	content := string(data)
	io.WriteString(file, content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
