package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Student Struct (Model)
type Student struct {
	ID    string `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
	AGE   int    `json:"age"`
}

// Init students var as a slice Student struct
var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// fmt.Println("params", params)
	for _, student := range students {
		if student.ID == params["id"] {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode("Student not found!")

}
func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)

	// fmt.Println("student", student)
	student.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	students = append(students, student)
	json.NewEncoder(w).Encode(student)

}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, student := range students {
		if student.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)

			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			// fmt.Println("student", student)

			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}

}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)

}

func main() {
	//Init Router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	students = append(students, Student{ID: "1", NAME: "ali", EMAIL: "a@gmail.com", AGE: 22})
	students = append(students, Student{ID: "2", NAME: "hamza", EMAIL: "h@gmail.com", AGE: 23})
	students = append(students, Student{ID: "3", NAME: "qadir", EMAIL: "q@gmail.com", AGE: 24})
	// fmt.Println(students)

	// Route Handlers / Endpoints
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/student/{id}", getStudent).Methods("GET")
	r.HandleFunc("/student", addStudent).Methods("POST")
	r.HandleFunc("/student/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}
