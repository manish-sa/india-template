package employeeservice

import (
	"context"

	employeerepository "github.com/manish-sa/india-template/internal/repository/employee"
)

type ServiceEmployee interface {
	// Create(context.Context, employeedto.GetEmployee) error
}

type Service struct {
	ctx     context.Context
	empRepo employeerepository.EmployeeRepository
}

func NewEmployeeService(ctx context.Context, empRepo employeerepository.EmployeeRepository) ServiceEmployee {
	return &Service{
		ctx:     ctx,
		empRepo: empRepo,
	}
}
