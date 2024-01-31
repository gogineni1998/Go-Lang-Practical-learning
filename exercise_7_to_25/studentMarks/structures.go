package student_results

type Student struct {
	Name        string
	DateOfBirth string
	School      string
}

type Marks struct {
	StudentDetails Student
	Maths          int64
	Physics        int64
	Chemistry      int64
}

type StudentTotal struct {
	Name  string
	Total int64
}
