package storage

import (
	"context"

	"ashish.com/m/pkg/models"
)

type EmployeeStore interface {
	Get(ctx context.Context, uid string) (*models.Employee, error)
	Update(ctx context.Context, uid string, updates models.Employee, fields ...string) error
	Create(ctx context.Context, employee models.Employee) (*models.Employee, error)
}

type PersonStore interface {
	Get(ctx context.Context, uid string) (*models.Person, error)
	Update(ctx context.Context, uid string, updates models.Person, fields ...string) error
	Create(ctx context.Context, person models.Person) (*models.Person, error)
}
