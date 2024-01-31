package main

import "fmt"

func main() {
	variable := Address{}
	variable.getAddress()
	fmt.Println(variable)
}

func (address *Address) getAddress() {
	address.AddressLine1 = "3929 Bailey Avenue"
	address.AddressLine2 = "Buffalo"
	address.City = "Amherst"
	address.County = "Erie"
	address.Phno = "7377770539"
	address.Pincode = 14226
	fmt.Println(address)
}

func (city City) getAddress(pincode int) {
	fmt.Println(pincode)
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	County       string
	Pincode      int
	Phno         string
}

type City struct {
}
