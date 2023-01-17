package service

import (
	"context"

	"algogrit.com/emp_server/entities"
)

type EmployeeService interface {
	Index(ctx context.Context) ([]entities.Employee, error)
	Create(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error)
}
