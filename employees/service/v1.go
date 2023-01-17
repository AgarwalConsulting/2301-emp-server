package service

import (
	"context"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

type v1 struct {
	repo repository.EmployeeRepository
}

func (svc *v1) Index(ctx context.Context) ([]entities.Employee, error) {
	return svc.repo.ListAll()
}
func (svc *v1) Create(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return &v1{repo}
}
