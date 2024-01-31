package main

import (
	calculations "exercise_till_25/performCalculations"
	student_results "exercise_till_25/studentMarks"
	"fmt"
)

func main() {
	studentsData := student_results.AddStudentDetails()
	data := calculations.Total(studentsData)
	fmt.Printf("%s", data)
}
