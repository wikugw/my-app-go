package handlers

import (
	"my-app/services"
	"net/http"

	"my-app/helper"

	"github.com/gin-gonic/gin"
)

// GetDepartmentsHandler godoc
// @Summary Get all departments
// @Description Retrieve a list of all departments
// @Tags departments
// @Produce json
// @Success 200 {object} types.Response
// @Failure 500 {object} types.Response
// @Router /departments [get]
func GetDepartmentsHandler(c *gin.Context) {
	depts, err := services.GetAllDepartments()
	if err != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch departments", nil)
		return
	}
	helper.JSONResponse(c, http.StatusOK, "success", "", depts)
}
