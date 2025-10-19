package handlers

import (
	"net/http"

	"my-app/helper"
	"my-app/services"

	employee "my-app/types/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new employee
// @Description Create employee with fullName, email, position, department
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body employee.CreateEmployeeRequest true "Employee info"
// @Success 201 {object} types.Response
// @Failure 400 {object} types.Response
// @Router /employees [post]
func CreateEmployeeHandler(c *gin.Context) {
	var req employee.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.JSONResponse(c, http.StatusBadRequest, "error", "Invalid request body: "+err.Error(), nil)
		return
	}

	employee, err := services.CreateEmployee(req)
	if err != nil {
		switch err {
		case services.ErrEmployeeExists:
			helper.JSONResponse(c, http.StatusConflict, "error", err.Error(), nil)
		default:
			helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to create employee", nil)
		}
		return
	}

	helper.JSONResponse(c, http.StatusCreated, "success", "", employee)
}

// @Summary Get employee by email
// @Description Get employee info by email query parameter
// @Tags employees
// @Produce json
// @Param email query string true "Email of the employee"
// @Success 200 {object} types.Response
// @Failure 400 {object} types.Response
// @Failure 404 {object} types.Response
// @Router /employees [get]
func GetEmployeeByEmailHandler(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		helper.JSONResponse(c, http.StatusBadRequest, "error", "Email query param is required", nil)
		return
	}

	employee, err := services.GetEmployeeByEmail(email)
	if err != nil {
		switch err {
		case services.ErrEmployeeNotFound:
			helper.JSONResponse(c, http.StatusNotFound, "error", err.Error(), nil)
		default:
			helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch employee", nil)
		}
		return
	}

	helper.JSONResponse(c, http.StatusOK, "success", "", employee)
}
