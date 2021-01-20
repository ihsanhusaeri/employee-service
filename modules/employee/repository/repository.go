package repository

import (
	"context"

	"github.com/ihsan-husaeri/employee-service/modules/employee/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type employeeRepository struct {
	database *mongo.Database
}

//EmployeeRepository interface
type EmployeeRepository interface {
	CreateEmployee(context.Context, *entity.Employee) (*entity.Employee, error)
	UpdateEmployee(context.Context, *entity.Employee) (*entity.Employee, error)
	GetAll(context.Context) (*[]entity.Employee, error)
}

//NewEmployeeRepository to create new instance
func NewEmployeeRepository(database *mongo.Database) EmployeeRepository {
	return &employeeRepository{database: database}
}
