package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.linkedin.com:3000/in/gdheeraj1998/?name=dheeraj&age=25"

func main() {
	fmt.Println("Handeling urls's")
	result, err := url.Parse(myurl)
	checkError(err)

	protocol := result.Scheme
	portNum := result.Port()
	url_path := result.Path
	qparams := result.Query()

	fmt.Println(protocol)
	fmt.Println(portNum)
	fmt.Println(url_path)

	for key, value := range qparams {
		fmt.Println(key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:   result.Scheme,
		Host:     result.Host,
		Path:     result.Path,
		RawQuery: result.RawQuery,
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
