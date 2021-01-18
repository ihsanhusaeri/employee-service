package handler

import (
	"github.com/ihsan-husaeri/employee-service/modules/employee/service"
	"github.com/labstack/echo/v4"
)

//EmployeeHandler struct
type employeeHandler struct {
	employeeService service.EmployeeService
}

//ApplyEmployeeHandler to create all employee route
func ApplyEmployeeHandler(e *echo.Echo, employeeService service.EmployeeService) {
	handler := &employeeHandler{
		employeeService: employeeService,
	}

	e.POST("/employee", handler.createEmployee)
	e.GET("/employee", handler.getAll)
}
