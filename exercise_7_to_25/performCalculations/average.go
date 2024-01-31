package calculations

import (
	"encoding/json"
	student_results "exercise_till_25/studentMarks"
)

func Total(data []byte) []byte {
	studentData := []student_results.Marks{}
	var totalMarks []student_results.StudentTotal
	err := json.Unmarshal(data, &studentData)
	if err != nil {
		panic(err)
	}
	for _, student := range studentData {
		totalMarks = append(totalMarks, student_results.StudentTotal{Name: student.StudentDetails.Name, Total: student.Maths + student.Chemistry + student.Physics})
	}
	totalData, err := json.Marshal(totalMarks)
	return totalData
}
