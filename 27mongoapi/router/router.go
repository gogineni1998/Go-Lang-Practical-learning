package router

import (
	"net/http"

	"github.com/gogineni1998/Go-Lang-Practical-learning/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/addmovie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/markwatched/{id}", controller.MarkMovieWatched).Methods("POST")
	r.HandleFunc("/delete/{id}", controller.DeleteOneMovie).Methods("POST")
	r.HandleFunc("/delete", controller.DeleteAllMovies).Methods("POST")
	r.HandleFunc("/getmovies", controller.GetAllMovies).Methods("GET")
	http.ListenAndServe(":4000", r)
}
