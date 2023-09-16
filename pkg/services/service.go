package services

import (
	"context"

	"ashish.com/m/pkg/models"
)

// EmployeeService is the base interface for building employee service.
type EmployeeService interface {
	GetEmployee(ctx context.Context, uid string) (*models.Employee, error)
	CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
}

// PersonService is the base interface for building Person service.
type PersonService interface {
	GetPerson(ctx context.Context, uid string) (*models.Person, error)
	UpdatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
	CreatePerson(ctx context.Context, person *models.Person) (*models.Person, error)
}
