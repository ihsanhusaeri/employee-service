package handler

import (
	"net/http"

	log "github.com/ihsan-husaeri/employee-service/common/logger"
	"github.com/ihsan-husaeri/employee-service/common/response"

	"github.com/ihsan-husaeri/employee-service/modules/employee/entity"
	"github.com/labstack/echo/v4"
)

func (emplHandler *employeeHandler) createEmployee(e echo.Context) error {
	var employee entity.Employee

	log.MakeLogEntry(e).Info("Incoming request: create new employee")
	ctx := e.Request().Context()

	if err := e.Bind(&employee); err != nil {
		log.MakeLogEntry(e).Error("error when binding data:", err)
		return e.JSON(response.NewResponse(http.StatusBadRequest, err, nil))

	}

	res, err := emplHandler.employeeService.CreateEmployee(ctx, &employee)

	if err != nil {
		return e.JSON(response.NewResponse(http.StatusInternalServerError, err, nil))
	}
	log.MakeLogEntry(e).Info("Employee Created", res.ID)
	return e.JSON(response.NewResponse(http.StatusCreated, "OK", res))
}

func (emplHandler *employeeHandler) getAll(e echo.Context) error {
	log.MakeLogEntry(e).Info("Incoming request: get all employee")
	ctx := e.Request().Context()

	res, err := emplHandler.employeeService.GetAll(ctx)

	if err != nil {
		return e.JSON(response.NewResponse(http.StatusInternalServerError, err, nil))
	}

	return e.JSON(response.NewResponse(http.StatusOK, "OK", res))
}
