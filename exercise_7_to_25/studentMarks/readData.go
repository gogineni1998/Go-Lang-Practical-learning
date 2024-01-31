package student_results

import (
	"os"
)

func readData() string {
	byteData, err := os.ReadFile("./data/input.txt")
	if err != nil {
		panic(err)
	}
	data := string(byteData)
	return data
}
