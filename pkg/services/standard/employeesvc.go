package standard

import (
	"context"

	errpkg "ashish.com/m/internal/errors"
	"ashish.com/m/pkg/models"
	storage "ashish.com/m/pkg/storage"
	log "github.com/sirupsen/logrus"
)

// EmployeeService is the standard implementation of employee service.
type EmployeeService struct {
	employeeStore storage.EmployeeStore
}

var (
	employeeIDEmptyErr  = "employee id is empty"
)

// NewEmployeeService creates a new instance of employee service.
func NewEmployeeService(s storage.EmployeeStore) *EmployeeService {
	return &EmployeeService{
		employeeStore: s,
	}
}

// GetEmployee returns the employee associated with the uid.
func (s *EmployeeService) GetEmployee(ctx context.Context, uid string) (*models.Employee, error) {
	fields := log.Fields{"employee.id": uid}

	// check uid
	if len(uid) == 0 {
		log.WithContext(ctx).WithFields(fields).Error(employeeIDEmptyErr)
		return nil, errpkg.NewEmptyError("employee id")
	}

	// get employee from data store
	employee, err := s.employeeStore.Get(ctx, uid)
	if err != nil {
		if _, ok := err.(*errpkg.RecordNotFoundError); ok {
			return nil, errpkg.NewResourceNotFoundError("employee", uid)
		}
		return nil, err
	}
	return employee, nil
}

// UpdateEmployeeError updates the employee error with the uid.
func (s *EmployeeService) CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	emp, err := s.employeeStore.Create(ctx, *employee)
	if err != nil {
		return nil, err
	}
	return emp, nil
}
