package service

import (
	"context"
	"errors"
	log "manyoption/payment-service/middleware/logger"

	// "time"

	"github.com/ihsan-husaeri/employee-service/modules/employee/entity"
)

func (emplService *employeeService) CreateEmployee(ctx context.Context, employee *entity.Employee) (*entity.Employee, error) {
	// ctx, cancel := context.WithTimeout(ctx, emplService.contextTimeout)
	// defer cancel()

	result, err := emplService.employeeRepo.CreateEmployee(emplService.ctx, employee)
	if err != nil {
		log.MakeLogEntry(nil).Panic(err)
		return nil, errors.New("Internal server error")
	}

	return result, nil
}

func (emplService *employeeService) GetAll(ctx context.Context) (*[]entity.Employee, error) {
	result, err := emplService.employeeRepo.GetAll(ctx)

	if err != nil {
		log.MakeLogEntry(nil).Panic(err)
		return nil, errors.New("Internal server error")
	}

	return result, nil
}
