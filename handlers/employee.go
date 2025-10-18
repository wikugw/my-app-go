package handlers

import (
	"net/http"

	"my-app/services"

	"github.com/gin-gonic/gin"

	"my-app/types"
	"my-app/types/services/employee"
)

func JSONResponse(c *gin.Context, code int, status, message string, data interface{}) {
	c.JSON(code, types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// @Summary Create a new employee
// @Description Create employee with fullName, email, position, department
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body services.CreateEmployeeRequest true "Employee info"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Router /employees [post]
func CreateEmployeeHandler(c *gin.Context) {
	var req employee.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", "Invalid request body: "+err.Error(), nil)
		return
	}

	employee, err := services.CreateEmployee(req)
	if err != nil {
		switch err {
		case services.ErrEmployeeExists:
			JSONResponse(c, http.StatusConflict, "error", err.Error(), nil)
		default:
			JSONResponse(c, http.StatusInternalServerError, "error", "Failed to create employee", nil)
		}
		return
	}

	JSONResponse(c, http.StatusCreated, "success", "", employee)
}

// @Summary Get employee by email
// @Description Get employee info by email query parameter
// @Tags employees
// @Produce json
// @Param email query string true "Email of the employee"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /employees [get]
func GetEmployeeByEmailHandler(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		JSONResponse(c, http.StatusBadRequest, "error", "Email query param is required", nil)
		return
	}

	employee, err := services.GetEmployeeByEmail(email)
	if err != nil {
		switch err {
		case services.ErrEmployeeNotFound:
			JSONResponse(c, http.StatusNotFound, "error", err.Error(), nil)
		default:
			JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch employee", nil)
		}
		return
	}

	JSONResponse(c, http.StatusOK, "success", "", employee)
}
