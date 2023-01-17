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

	var empRepo = repository.NewInMem()
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empSvc)

	empHandler.SetupRoutes(r)

	log.Info("Starting the server on port: 8000...")
	// err := http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))

	log.Fatal(err)
}
