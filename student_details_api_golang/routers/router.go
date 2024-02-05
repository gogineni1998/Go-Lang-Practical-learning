package router

import (
	"net/http"

	"github.com/gogineni1998/student_details_api_golang/controllers"
	"github.com/gorilla/mux"
)

var r *mux.Router

func Router() {
	r = mux.NewRouter()
	r.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/delete/{id}", controllers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/update/{id}", controllers.UpdateStudent).Methods("PATCH")
	http.ListenAndServe(":3000", r)
}
