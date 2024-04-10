package employeerepository

import (
	"context"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	// Create(ctx context.Context) (uint64, error)
}

type Repository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewEmployeeRepository(ctx context.Context, db *gorm.DB) EmployeeRepository {
	return &Repository{
		ctx: ctx,
		db:  db,
	}
}
