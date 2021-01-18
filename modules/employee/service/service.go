package service

import (
	"context"

	"github.com/ihsan-husaeri/employee-service/modules/employee/entity"
	"github.com/ihsan-husaeri/employee-service/modules/employee/repository"
)

//EmployeeService interface
type EmployeeService interface {
	CreateEmployee(context.Context, *entity.Employee) (*entity.Employee, error)
	GetAll(context.Context) (*[]entity.Employee, error)
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
	ctx          context.Context
}

//NewEmployeeService to create new instance
func NewEmployeeService(ctx context.Context, employeeRepository repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepo: employeeRepository,
		ctx:          ctx,
	}
}
