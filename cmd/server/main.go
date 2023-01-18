package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/employees/service"
)

func envOrDefault(key, dfltVal string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return dfltVal
	}

	return val
}

var port string = "8000"

func init() {
	port = envOrDefault("PORT", port)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	// var empRepo = repository.NewInMem()
	var empRepo = repository.NewSQL("postgres", "postgresql://localhost:5432/emp-demo?sslmode=disable")
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empSvc)

	empHandler.SetupRoutes(r)

	log.Infof("Starting the server on port: %s...", port)
	// err := http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))

	log.Fatal(err)
}
