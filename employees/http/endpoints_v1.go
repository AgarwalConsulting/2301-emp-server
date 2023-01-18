package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/emp_server/entities"

	"github.com/go-playground/validator/v10"
)

func (h *EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	emps, err := h.svcV1.Index(req.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func (h *EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	validate := validator.New()
	errs := validate.Struct(newEmp)

	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errs)
		return
	}

	savedEmp, err := h.svcV1.Create(req.Context(), newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(savedEmp)
}
