package repository

import (
	"sync"

	"algogrit.com/emp_server/entities"
)

type inmem struct {
	employees []entities.Employee
	mut       *sync.RWMutex
}

func (repo *inmem) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock() // Exclusive for Writing (allows multiple readers)
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inmem) Save(newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock() // Exclusive for Writing & Reading (doesn't allow anyone else)
	defer repo.mut.Unlock()

	newEmp.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Prathyash", "Cloud", 10001},
		{3, "Anita", "SRE", 20001},
	}

	return &inmem{employees, &sync.RWMutex{}}
}
