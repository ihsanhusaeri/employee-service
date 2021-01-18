package repository

import (
	"context"
	"errors"

	"github.com/ihsan-husaeri/employee-service/common"
	"github.com/ihsan-husaeri/employee-service/modules/employee/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (emplRepository *employeeRepository) CreateEmployee(ctx context.Context, empl *entity.Employee) (*entity.Employee, error) {
	result, err := emplRepository.database.Collection(common.EmployeeCollection).InsertOne(ctx, empl)
	// fmt.Println(result)
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("Failed convert to primitive object ID")
	}
	empl.ID = &id
	return empl, err
}

func (emplRepository *employeeRepository) GetAll(ctx context.Context) (*[]entity.Employee, error) {
	result, err := emplRepository.database.Collection(common.EmployeeCollection).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var empls []entity.Employee

	if err = result.All(ctx, &empls); err != nil {
		return nil, err
	}

	return &empls, nil
}
