package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonObject := encodeJson()
	address := decodeJson(jsonObject)
	fmt.Printf("%+v", address)
}

func decodeJson(jsonObject []byte) []Address {
	var address []Address
	err := json.Unmarshal(jsonObject, &address)
	if err != nil {
		panic(err)
	}
	return address
}

func encodeJson() []byte {
	requestBody := []Address{
		{"3929 Bailey Avenue", "Buffalo", "Erie", "USA", 14226, "775-10-7746"},
		{"40 Englewood Avenue", "Buffalo", "Erie", "USA", 14214, "775-10-7746"},
	}
	jsonObj, err := json.MarshalIndent(requestBody, "", "\t")

	if err != nil {
		panic(err)
	}

	return jsonObj
}

type Address struct {
	Street  string `json:"streetName"`
	City    string `json:"cityName"`
	County  string `json:"countyName,omitempty"`
	Country string `json:"countryName"`
	Pin     int16  `json:"pinNumber"`
	Ssn     string `json:"-"`
}
