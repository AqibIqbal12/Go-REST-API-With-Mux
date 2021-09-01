package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Student Struct (Model)
type Student struct {
	ID    string `json:id`
	NAME  string `json:name`
	EMAIL string `json:email`
	AGE   int    `json:age`
}

func main() {
	//Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("api/students", getStudents).Methods("GET")
	r.HandleFunc("api/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("api/students", addStudent).Methods("POST")
	r.HandleFunc("api/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("api/students/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}
