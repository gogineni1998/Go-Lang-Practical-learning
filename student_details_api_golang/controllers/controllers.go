package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gogineni1998/student_details_api_golang/dbhelpers"
	"github.com/gogineni1998/student_details_api_golang/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var student *models.Student

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)

	errorHandler(err)
	err = json.Unmarshal(data, &student)
	errorHandler(err)

	response := dbhelpers.InsertRecord(student)
	w.Write([]byte(response.InsertedID.(primitive.ObjectID).String()))
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	id := mux.Vars(r)
	errorHandler(err)
	err = json.Unmarshal(data, &student)
	errorHandler(err)
	response := dbhelpers.UpdateRecord(id["id"], student)
	w.Write([]byte(response))
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	response := dbhelpers.DeleteRecord(id["id"])
	w.Write([]byte(strconv.Itoa(int(response))))
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
