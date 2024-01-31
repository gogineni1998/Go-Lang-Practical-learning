package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Creating Json Data")
	peopleData := []person{
		{"Dheeraj Sai Gogineni", 25, "3929 Bailey Avenue", 14226},
	}

	jsonData, err := json.MarshalIndent(peopleData, ".", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", jsonData)
}

type person struct {
	Name    string `json:"personName"`
	Age     int8   `json:"personAge"`
	Address string `json:"houseAddress,omitempty"`
	Pin     int16  `json:"-"`
}
