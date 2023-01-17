package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"algogrit.com/emp_server/entities"
)

var employees = []entities.Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Prathyash", "Cloud", 10001},
	{3, "Anita", "SRE", 20001},
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	handler := func(w http.ResponseWriter, req *http.Request) {
// 		begin := time.Now()

// 		next.ServeHTTP(w, req)

// 		log.Infof("%s %s took %s\n", req.Method, req.URL, time.Since(begin))
// 	}

// 	return http.HandlerFunc(handler)
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	log.Info("Starting the server on port: 8000...")
	// err := http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))

	log.Fatal(err)
}
