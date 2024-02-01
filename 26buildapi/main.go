package main

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/cources", getAllCourses).Methods("GET")
	r.HandleFunc("/cources/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/create", createOneCourse).Methods("POST")
	r.HandleFunc("/update", updateOneCourse).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteOneCourse).Methods("POST")
	addSeeds()
	http.ListenAndServe(":4000", r)
}

func addSeeds() {
	courses = append(courses, Course{"1001", "AWS", "$200", &Author{"Dheeraj", "http://dheeraj.com"}})
	courses = append(courses, Course{"1002", "GCP", "$150", &Author{"Satya Neeraj Gogineni", "http://satya.com"}})
	courses = append(courses, Course{"1003", "AZURE", "$300", &Author{"Dheeraj", "http://dheeraj.com"}})
	courses = append(courses, Course{"1004", "ALICLOUD", "$50", &Author{"Satya Neeraj Gogineni", "http://satya.com"}})
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	output, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course

	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		panic(err)
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Add a Course Name")
		return
	}

	rand.New(rand.NewSource(time.Now().UnixMicro()))

	index := rand.Intn(100)

	course.CourseId = strconv.Itoa(index)

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var updatedCource Course
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &updatedCource)
	if err != nil {
		panic(err)
	}
	for ind, course := range courses {
		if course.CourseId == updatedCource.CourseId {
			courses = append(courses[:ind], courses[ind+1:]...)
			courses = append(courses, updatedCource)
			break
		}
	}

	json.NewEncoder(w).Encode(&courses)

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	for ind, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:ind], courses[ind+1:]...)
			json.NewEncoder(w).Encode(&courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Cource Not Found")
}
