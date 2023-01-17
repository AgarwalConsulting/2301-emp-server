package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"` // Struct Tags
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"id": %d, "name": "%s", "speciality": "%s"}`, e.ID, e.Name, e.Department)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Prathyash", "Cloud", 10001},
	{3, "Anita", "SRE", 20001},
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&newEmp)

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		EmployeeCreateHandler(w, req)
	} else {
		EmployeesIndexHandler(w, req)
	}
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesHandler)

	// r.HandleFunc("/employees", EmployeesIndexHandler)
	// r.HandleFunc("/employees", EmployeeCreateHandler)

	err := http.ListenAndServe("localhost:8000", r)

	log.Fatal(err)
}
