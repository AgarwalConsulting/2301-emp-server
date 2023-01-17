package repository

import "algogrit.com/emp_server/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
