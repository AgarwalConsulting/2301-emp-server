package service

import (
	"context"

	"algogrit.com/emp_server/entities"
)

//go:generate mockgen -source=$GOFILE -destination ./mock_$GOPACKAGE.go -package $GOPACKAGE

type EmployeeService interface {
	Index(ctx context.Context) ([]entities.Employee, error)
	Create(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error)
}
