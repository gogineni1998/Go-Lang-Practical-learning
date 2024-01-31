package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	myMap := make(map[string]string)

	myMap["GDS"] = "Dheeraj Sai Gogineni"
	myMap["SGN"] = "Satya Neeraj Gogineni"
	myMap["VG"] = "Vijayababu Gogineni"
	myMap["PG"] = "Praveena Gogineni"
	myMap["DD"] = "Daridram"

	fmt.Println(myMap)

	delete(myMap, "DD")

	fmt.Println(myMap)

	fmt.Println("Looping through the map")

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
