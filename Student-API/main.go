package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)


type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Course string `json:"course"`
}

var students = []Student{
	{1, "Rahat", "BCA"},
	{2, "Alina", "BCA"},
	{3, "Riya", "BCA"},
}

// GET: Fetch one student
func get_student(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/student-data/")

	num, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	for _, s := range students {
		if s.Id == num {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
			return
		}
	}

	http.Error(w, "Student not found", 404)
}

// GET: Fetch all students
func get_all_students(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// POST: Add new student
func add_student(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var newStudent Student

	err := json.NewDecoder(r.Body).Decode(&newStudent)

	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	students = append(students, newStudent)
		
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newStudent)
}

// Main function
func main() {

	http.HandleFunc("/student-data/", get_student)
	http.HandleFunc("/students", get_all_students)
	http.HandleFunc("/add-student", add_student)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
