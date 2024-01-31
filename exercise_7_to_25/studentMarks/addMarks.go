package student_results

import (
	"encoding/json"
	"strconv"
	"strings"
)

func AddStudentDetails() []byte {
	data := readData()
	var marks []Marks
	studentDetials := strings.Split(data, "\n")
	for _, studentDetial := range studentDetials {
		studentMarks := strings.Split(studentDetial, " ")
		student := Student{studentMarks[0], studentMarks[1], studentMarks[2]}

		maths, err := strconv.ParseInt(studentMarks[3], 0, 64)
		if err != nil {
			panic(err)
		}

		physics, err := strconv.ParseInt(studentMarks[4], 0, 64)
		if err != nil {
			panic(err)
		}

		chemistry, err := strconv.ParseInt(strings.TrimSpace(studentMarks[5]), 0, 64)
		if err != nil {
			panic(err)
		}
		marks = append(marks, Marks{student, maths, physics, chemistry})
	}
	jsonData, err := json.MarshalIndent(marks, "", "\t")
	if err != nil {
		panic(err)
	}
	return jsonData
}
