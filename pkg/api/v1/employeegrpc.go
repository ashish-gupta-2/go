package v1

import (
	"context"
	"strconv"

	"ashish.com/m/internal/utils"
	v1 "ashish.com/m/pb/ashish.com/v1"
	"ashish.com/m/pkg/models"
	"ashish.com/m/pkg/services"
)

// EmployeeGRPCService is the grpc service implementation of employee service.
type EmployeeGRPCService struct {
	employeeSvc services.EmployeeService
	v1.UnimplementedEmployeeServiceServer
}

// NewEmployeeGRPCService creates a new instance of employee grpc service.
func NewEmployeeGRPCService(employeeSvc services.EmployeeService) *EmployeeGRPCService {
	return &EmployeeGRPCService{
		employeeSvc: employeeSvc,
	}
}

// CreateEmployee creates a new Employee using specified request.
func (s *EmployeeGRPCService) CreateEmployee(ctx context.Context, req *v1.CreateEmployeeRequest) (*v1.CreateEmployeeResponse, error) {
	employee := &models.Employee{
		Name:   req.Employee.Name,
		Salary: strconv.FormatFloat(req.Employee.Salary, 'g', 5, 64),
	}

	// call standard employee service.
	employeeCreated, err := s.employeeSvc.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, utils.MapToGRPCStatus(err).Err()
	}
	return &v1.CreateEmployeeResponse{
		Employee: toGRPCEmployee(*employeeCreated),
	}, nil
}

// GetEmployee returns the employee associated with the resource name.
func (s *EmployeeGRPCService) GetEmployee(ctx context.Context, req *v1.GetEmployeeRequest) (*v1.GetEmployeeResponse, error) {
	// call standard employee service
	employee, err := s.employeeSvc.GetEmployee(ctx, req.Id)
	if err != nil {
		return nil, utils.MapToGRPCStatus(err).Err()
	}
	// return fetched employee
	return &v1.GetEmployeeResponse{
		Employee: toGRPCEmployee(*employee),
	}, nil
}

func toGRPCEmployee(employee models.Employee) *v1.Employee {
	salaryInFloat, _ := strconv.ParseFloat(employee.Salary, 64)
	// return fetched employee
	return &v1.Employee{
		Id:     employee.UID.String(),
		Name:   employee.Name,
		Salary: salaryInFloat,
	}
}
